package sender

import (
	"github.com/Shopify/sarama"
	"github.com/ozonmp/com-message-api/internal/config"
	"github.com/ozonmp/com-message-api/internal/model"
	pb "github.com/ozonmp/com-message-api/pkg/com-message-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//KafkaSender entity for sending events for kafka
type KafkaSender struct {
	producer sarama.SyncProducer
	topic    string
}

//NewKafkaSender constructor for kafka sender
func NewKafkaSender(config config.Kafka) (*KafkaSender, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Partitioner = sarama.NewRandomPartitioner
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(config.Brokers, saramaConfig)
	if err != nil {
		return nil, err
	}

	sender := KafkaSender{
		producer: producer,
		topic:    config.Topic,
	}

	return &sender, err
}

//Send send message to kafka
func (k KafkaSender) Send(messageEvent *model.MessageEvent) error {
	payload := messageEvent.Payload
	pbMessage := &pb.Message{}
	err := protojson.Unmarshal([]byte(payload), pbMessage)
	if err != nil {
		log.Error().Err(err).Msg("error appeared during unmarshaling payload")
	}

	pbMessageEvent := &pb.MessageEvent{
		Id:        messageEvent.ID,
		MessageId: messageEvent.MessageID,
		Type:      messageEvent.TypeDb,
		Status:    messageEvent.Status.String(),
		Payload:   pbMessage,
		Updated:   timestamppb.New(messageEvent.Updated.Time),
	}

	newMessage, _ := protojson.Marshal(pbMessageEvent)
	msg := &sarama.ProducerMessage{
		Topic:     k.topic,
		Partition: -1,
		Value:     sarama.ByteEncoder(newMessage),
	}

	_, _, err = k.producer.SendMessage(msg)

	return err
}
