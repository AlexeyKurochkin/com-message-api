package api

import (
	"context"

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
		Value: &pb.Message{
			Id:  message.ID,
			Foo: message.Foo,
		},
	}, nil
}
