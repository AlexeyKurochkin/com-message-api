package message

import (
	"context"
	"fmt"
	"github.com/ozonmp/com-message-api/internal/config"
	"github.com/ozonmp/com-message-api/internal/model"
	proto "github.com/ozonmp/com-message-api/pkg/com-message-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//Service for bot
type Service struct {
	client proto.ComMessageApiServiceClient
}

//NewMessageService constructor for Service
func NewMessageService(cfg *config.Config) *Service {
	connection, error := grpc.Dial(fmt.Sprintf("%v:%v", cfg.Grpc.Host, cfg.Grpc.Port), grpc.WithInsecure())
	if error != nil {
		log.Error().Err(error).Msg("Error on creating bot message service")
	}

	client := proto.NewComMessageApiServiceClient(connection)
	return &Service{
		client: client,
	}
}

//Describe message by id
func (d *Service) Describe(messageID uint64) (*model.Message, error) {
	ctx := context.Background()
	response, error := d.client.DescribeMessageV1(ctx, &proto.DescribeMessageV1Request{MessageId: messageID})
	if error != nil {
		return nil, error
	}

	return convertPbModelToMessage(response.Value), nil
}

func convertPbModelToMessage(pbMessage *proto.Message) *model.Message {
	return &model.Message{
		ID:       pbMessage.GetId(),
		From:     pbMessage.GetFrom(),
		To:       pbMessage.GetTo(),
		Text:     pbMessage.GetText(),
		Datetime: pbMessage.GetDatetime().AsTime(),
	}
}

//List messages
func (d *Service) List(cursor uint64, limit uint64) ([]*model.Message, error) {
	ctx := context.Background()
	response, error := d.client.ListMessageV1(ctx, &proto.ListMessageV1Request{})
	if error != nil {
		return nil, error
	}

	messages := make([]*model.Message, 0, len(response.Value))
	for _, message := range response.Value {
		messages = append(messages, convertPbModelToMessage(message))
	}
	return messages, nil
}

//Create new message
func (d *Service) Create(message *model.Message) (uint64, error) {
	ctx := context.Background()
	response, error := d.client.CreateMessageV1(ctx, &proto.CreateMessageV1Request{
		From:     message.From,
		To:       message.To,
		Text:     message.Text,
		Datetime: timestamppb.New(message.Datetime),
	})
	if error != nil {
		return 0, error
	}

	return response.Value.GetId(), nil
}

//Update message
func (d *Service) Update(messageID uint64, message *model.Message) error {
	ctx := context.Background()
	_, error := d.client.UpdateMessageV1(ctx, &proto.UpdateMessageV1Request{
		MessageId: messageID,
		From:      message.From,
		To:        message.To,
		Text:      message.Text,
		Datetime:  timestamppb.New(message.Datetime),
	})
	if error != nil {
		return error
	}

	return nil
}

//Remove message by id
func (d *Service) Remove(messageID uint64) (bool, error) {
	ctx := context.Background()
	_, error := d.client.RemoveMessageV1(ctx, &proto.RemoveMessageV1Request{MessageId: messageID})
	if error != nil {
		return false, error
	}

	return true, nil
}
