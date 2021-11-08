package producer

import (
	"context"
	"github.com/gammazero/workerpool"
	"github.com/ozonmp/com-message-api/internal/app/repo"
	"github.com/ozonmp/com-message-api/internal/app/sender"
	"github.com/ozonmp/com-message-api/internal/model"
	"log"
	"sync"
)

//Producer interface
type Producer interface {
	Start(ctx context.Context)
	Close()
}

type producer struct {
	producerCount uint64
	sender        sender.EventSender
	events        <-chan model.MessageEvent
	workerPool    *workerpool.WorkerPool
	repo          repo.EventRepo
	wg            *sync.WaitGroup
}

// Config for Producer
type Config struct {
	ProducerCount uint64
	Sender        sender.EventSender
	Events        <-chan model.MessageEvent
	WorkerPool    *workerpool.WorkerPool
	Repo          repo.EventRepo
}

// NewKafkaProducer constructor for producer
func NewKafkaProducer(config Config) Producer {
	wg := &sync.WaitGroup{}
	return &producer{
		producerCount: config.ProducerCount,
		sender:        config.Sender,
		events:        config.Events,
		workerPool:    config.WorkerPool,
		repo:          config.Repo,
		wg:            wg,
	}
}

func (p *producer) Start(ctx context.Context) {
	for i := uint64(0); i < p.producerCount; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case event := <-p.events:
					processEvent(p, event)
				case <-ctx.Done():
					return
				}
			}
		}()
	}
}

func (p *producer) Close() {
	p.wg.Wait()
}

func processEvent(p *producer, event model.MessageEvent) {
	if error := p.sender.Send(&event); error != nil {
		processUnsuccessfulSending(p, event)
	} else {
		processSuccessfulSending(p, event)
	}
}

func processUnsuccessfulSending(p *producer, event model.MessageEvent) {
	log.Printf("Error sending event to kafka:\n%v", event)
	tryUnlockEvent(p, event)
}

func tryUnlockEvent(p *producer, event model.MessageEvent) {
	p.workerPool.Submit(func() {
		log.Println("Trying unlock event...")
		if err := p.repo.Unlock([]uint64{event.ID}); err != nil {
			log.Printf("Error unlocking event:\n%v", event)
		} else {
			log.Printf("Successfully unlocked event:\n%v", event)
		}
	})
}

func processSuccessfulSending(p *producer, event model.MessageEvent) {
	log.Printf("Successfully sent event to kafka:\n%v", event)
	tryRemoveEvent(p, event)
}

func tryRemoveEvent(p *producer, event model.MessageEvent) {
	p.workerPool.Submit(func() {
		if err := p.repo.Remove([]uint64{event.ID}); err != nil {
			log.Printf("Error removing event:\n%v", event)
		} else {
			log.Printf("Successfully removed event:\n%v", event)
		}
	})
}
