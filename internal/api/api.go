package api

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/ozonmp/com-message-api/internal/logging"
	"github.com/ozonmp/com-message-api/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	eventRepo "github.com/ozonmp/com-message-api/internal/app/repo"
	"github.com/ozonmp/com-message-api/internal/repo"

	pb "github.com/ozonmp/com-message-api/pkg/com-message-api"
)

var (
	totalMessageNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "com_message_api_message_not_found_total",
		Help: "Total number of messages that were not found",
	})
)

type messageAPI struct {
	pb.UnimplementedComMessageApiServiceServer
	repo      repo.Repo
	eventRepo eventRepo.EventRepo
}

// NewMessageAPI returns api of com-message-api service
func NewMessageAPI(r repo.Repo, er eventRepo.EventRepo) pb.ComMessageApiServiceServer {
	return &messageAPI{repo: r, eventRepo: er}
}

func (o *messageAPI) CreateMessageV1(
	ctx context.Context,
	req *pb.CreateMessageV1Request,
) (*pb.CreateMessageV1Response, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateMessageV1")
	defer span.Finish()

	log := logging.GetLoggerAssociatedWithCtx(ctx, "api", "CreateMessageV1")

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	message := &model.Message{
		From:     req.GetFrom(),
		To:       req.GetTo(),
		Text:     req.GetText(),
		Datetime: req.GetDatetime().AsTime(),
	}

	newID, err := o.repo.CreateMessage(ctx, message)
	if err != nil {
		log.Error().Err(err).Msg("failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	messageEvent := model.MessageEvent{
		MessageId: newID,
		TypeDb:    model.Created.String(),
		Status:    model.New,
		Type:      model.Created,
		Entity:    message,
	}
	err = o.eventRepo.Add(messageEvent)
	if err != nil {
		log.Error().Err(err).Msg("Adding event failed")
	}

	log.Debug().Msg("success")

	message.ID = newID
	return &pb.CreateMessageV1Response{
		Value: convertMessageToPbModel(message),
	}, nil
}

func (o *messageAPI) DescribeMessageV1(
	ctx context.Context,
	req *pb.DescribeMessageV1Request,
) (*pb.DescribeMessageV1Response, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "DescribeMessageV1")
	defer span.Finish()

	log := logging.GetLoggerAssociatedWithCtx(ctx, "api", "DescribeMessageV1")

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	message, err := o.repo.DescribeMessage(ctx, req.MessageId)
	if err != nil {
		log.Error().Err(err).Msg("failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if message == nil {
		log.Debug().Uint64("messageId", req.MessageId).Msg("message not found")
		totalMessageNotFound.Inc()

		return nil, status.Error(codes.NotFound, "message not found")
	}

	log.Debug().Msg("success")

	return &pb.DescribeMessageV1Response{
		Value: convertMessageToPbModel(message),
	}, nil
}

func (o *messageAPI) ListMessageV1(
	ctx context.Context,
	req *pb.ListMessageV1Request,
) (*pb.ListMessageV1Response, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "ListMessageV1")
	defer span.Finish()

	log := logging.GetLoggerAssociatedWithCtx(ctx, "api", "ListMessageV1")

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	messages, err := o.repo.ListMessage(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if messages == nil {
		log.Debug().Msg("messages not found")
		totalMessageNotFound.Inc()

		return nil, status.Error(codes.NotFound, "message not found")
	}

	log.Debug().Msg("success")

	pbMessages := make([]*pb.Message, 0, len(messages))
	for _, message := range messages {
		currentMessage := message
		pbMessages = append(pbMessages, convertMessageToPbModel(&currentMessage))
	}

	return &pb.ListMessageV1Response{
		Value: pbMessages,
	}, nil
}

func convertMessageToPbModel(message *model.Message) *pb.Message {
	return &pb.Message{
		Id:       message.ID,
		From:     message.From,
		To:       message.To,
		Text:     message.Text,
		Datetime: timestamppb.New(message.Datetime),
	}
}

func (o *messageAPI) RemoveMessageV1(
	ctx context.Context,
	req *pb.RemoveMessageV1Request,
) (*pb.RemoveMessageV1Response, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "RemoveMessageV1")
	defer span.Finish()

	log := logging.GetLoggerAssociatedWithCtx(ctx, "api", "RemoveMessageV1")

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	messageId := req.GetMessageId()
	success, err := o.repo.RemoveMessage(ctx, messageId)
	if err != nil {
		log.Error().Err(err).Msg("failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if !success {
		log.Debug().Msg("message not found")
		totalMessageNotFound.Inc()

		return nil, status.Error(codes.NotFound, "message not found")
	}

	//todo check for error due to nil entity
	messageEvent := model.MessageEvent{
		MessageId: messageId,
		TypeDb:    model.Removed.String(),
		Status:    model.New,
		Type:      model.Removed,
	}
	err = o.eventRepo.Add(messageEvent)
	if err != nil {
		log.Error().Err(err).Msg("Adding event failed")
	}

	log.Debug().Msg("success")

	return &pb.RemoveMessageV1Response{
		Result: success,
	}, nil
}
