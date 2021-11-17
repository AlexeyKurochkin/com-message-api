package consumer

import (
	"context"
	"github.com/ozonmp/com-message-api/internal/app/repo"
	"github.com/ozonmp/com-message-api/internal/metrics"
	"github.com/ozonmp/com-message-api/internal/model"
	"log"
	"sync"
	"time"
)

//Consumer interface
type Consumer interface {
	Start(ctx context.Context)
	Close()
}

type consumer struct {
	consumerCount uint64
	batchSize     uint64
	timeout       time.Duration
	events        chan<- model.MessageEvent
	repo          repo.EventRepo
	wg            *sync.WaitGroup
}

//Config for consumer
type Config struct {
	ConsumerCount uint64
	BatchSize     uint64
	Timeout       time.Duration
	Events        chan<- model.MessageEvent
	Repo          repo.EventRepo
}

//NewDbConsumer constructor for consumer
func NewDbConsumer(config Config) Consumer {
	wg := &sync.WaitGroup{}
	return &consumer{
		consumerCount: config.ConsumerCount,
		batchSize:     config.BatchSize,
		timeout:       config.Timeout,
		events:        config.Events,
		repo:          config.Repo,
		wg:            wg,
	}
}

func (c *consumer) Start(ctx context.Context) {
	for i := uint64(0); i < c.consumerCount; i++ {
		c.wg.Add(1)

		go func() {
			defer c.wg.Done()
			ticker := time.NewTicker(c.timeout)
			for {
				select {
				case <-ticker.C:
					events, err := c.repo.Lock(c.batchSize)
					if err != nil {
						log.Printf("Error locking events: %v", err)
						continue
					}
					for _, event := range events {
						c.events <- event
						metrics.HandledEventsTotal.Add(1)
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}
}

func (c *consumer) Close() {
	close(c.events)
	c.wg.Wait()
}
