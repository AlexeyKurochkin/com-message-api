package api

import (
	"context"
	"github.com/ozonmp/com-message-api/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	repo repo.Repo
}

// NewMessageAPI returns api of com-message-api service
func NewMessageAPI(r repo.Repo) pb.ComMessageApiServiceServer {
	return &messageAPI{repo: r}
}

func (o *messageAPI) CreateMessageV1(
	ctx context.Context,
	req *pb.CreateMessageV1Request,
) (*pb.CreateMessageV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("CreateMessageV1 - invalid argument")

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
		log.Error().Err(err).Msg("CreateMessageV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("CreateMessageV1 - success")

	message.ID = newID
	return &pb.CreateMessageV1Response{
		Value: convertMessageToPbModel(message),
	}, nil
}

func (o *messageAPI) DescribeMessageV1(
	ctx context.Context,
	req *pb.DescribeMessageV1Request,
) (*pb.DescribeMessageV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeMessageV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	message, err := o.repo.DescribeMessage(ctx, req.MessageId)
	if err != nil {
		log.Error().Err(err).Msg("DescribeMessageV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if message == nil {
		log.Debug().Uint64("messageId", req.MessageId).Msg("message not found")
		totalMessageNotFound.Inc()

		return nil, status.Error(codes.NotFound, "message not found")
	}

	log.Debug().Msg("DescribeMessageV1 - success")

	return &pb.DescribeMessageV1Response{
		Value: convertMessageToPbModel(message),
	}, nil
}

func (o *messageAPI) ListMessageV1(
	ctx context.Context,
	req *pb.ListMessageV1Request,
) (*pb.ListMessageV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("ListMessageV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	messages, err := o.repo.ListMessage(ctx)
	if err != nil {
		log.Error().Err(err).Msg("ListMessageV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if messages == nil {
		log.Debug().Msg("messages not found")
		totalMessageNotFound.Inc()

		return nil, status.Error(codes.NotFound, "message not found")
	}

	log.Debug().Msg("ListMessageV1 - success")

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

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("RemoveMessageV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	success, err := o.repo.RemoveMessage(ctx, req.GetMessageId())
	if err != nil {
		log.Error().Err(err).Msg("RemoveMessageV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if !success {
		log.Debug().Msg("message not found")
		totalMessageNotFound.Inc()

		return nil, status.Error(codes.NotFound, "message not found")
	}

	log.Debug().Msg("RemoveMessageV1 - success")

	return &pb.RemoveMessageV1Response{
		Result: success,
	}, nil
}
