package sender

import (
	"github.com/Shopify/sarama"
	"github.com/ozonmp/com-message-api/internal/config"
	"github.com/ozonmp/com-message-api/internal/model"
	"github.com/rs/zerolog/log"
)

type KafkaSenderConfig struct {
	brokers []string
	topic   string
}

type KafkaSender struct {
	producer sarama.SyncProducer
	topic    string
}

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

func (k KafkaSender) Send(messageEvent *model.MessageEvent) error {
	msg := &sarama.ProducerMessage{
		Topic:     k.topic,
		Partition: -1,
		Value:     sarama.ByteEncoder(messageEvent.Payload),
	}

	log.Debug().Msgf("Sending message to kafka\nMessageId:%v;From:%v", messageEvent.Entity.ID, messageEvent.Entity.From)
	_, _, err := k.producer.SendMessage(msg)

	return err
}
