package retranslator

import (
	"context"
	"github.com/gammazero/workerpool"
	"github.com/ozonmp/com-message-api/internal/app/consumer"
	"github.com/ozonmp/com-message-api/internal/app/producer"
	"github.com/ozonmp/com-message-api/internal/app/repo"
	"github.com/ozonmp/com-message-api/internal/app/sender"
	"github.com/ozonmp/com-message-api/internal/model"
	"time"
)

type Retranslator interface {
	Start(ctx context.Context)
	Close()
}

type Config struct {
	ConsumerCount  uint64
	ConsumeSize    uint64
	ConsumeTimeout time.Duration
	ProducerCount  uint64
	WorkerCount    int
	ChannelSize    uint64
	Repo           repo.EventRepo
	Sender         sender.EventSender
}

type retranslator struct {
	events     chan model.MessageEvent
	consumer   consumer.Consumer
	producer   producer.Producer
	workerPool *workerpool.WorkerPool
}

func NewRetranslator(cfg Config) Retranslator {
	events := make(chan model.MessageEvent, cfg.ChannelSize)
	workerPool := workerpool.New(cfg.WorkerCount)
	consumerConfig := getConsumerConfig(cfg, events)
	consumer := consumer.NewDbConsumer(consumerConfig)
	producerConfig := getProducerConfig(cfg, events, workerPool)
	producer := producer.NewKafkaProducer(producerConfig)

	return &retranslator{
		events:     events,
		consumer:   consumer,
		producer:   producer,
		workerPool: workerPool,
	}
}

func getProducerConfig(cfg Config, events chan model.MessageEvent, workerPool *workerpool.WorkerPool) producer.Config {
	producerConfig := producer.Config{
		ProducerCount: cfg.ProducerCount,
		Sender:        cfg.Sender,
		Events:        events,
		WorkerPool:    workerPool,
		Repo:          cfg.Repo,
	}
	return producerConfig
}

func getConsumerConfig(cfg Config, events chan model.MessageEvent) consumer.Config {
	consumerConfig := consumer.Config{
		ConsumerCount: cfg.ConsumerCount,
		BatchSize:     cfg.ConsumeSize,
		Timeout:       cfg.ConsumeTimeout,
		Events:        events,
		Repo:          cfg.Repo,
	}
	return consumerConfig
}

func (r *retranslator) Start(ctx context.Context) {
	r.producer.Start(ctx)
	r.consumer.Start(ctx)
}

func (r *retranslator) Close() {
	r.consumer.Close()
	r.producer.Close()
	r.workerPool.Stop()
}
