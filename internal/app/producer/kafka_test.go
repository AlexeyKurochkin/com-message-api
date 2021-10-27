package producer

import (
	"github.com/ozonmp/com-message-api/internal/mocks"
	"github.com/ozonmp/com-message-api/internal/model"
	"context"
	"errors"
	"github.com/gammazero/workerpool"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var testEvent = model.MessageEvent{
	ID:     1,
	Type:   model.Created,
	Status: model.Deferred,
	Entity: &model.Message{},
}

func TestProducer_NewKafkaProducer(t *testing.T){
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	events := make(chan model.MessageEvent, 1)
	events <- testEvent

	workerPool := workerpool.New(1)
	config := Config{
		1,
		sender,
		events,
		workerPool,
		repo,
	}
	actualResult := NewKafkaProducer(config)
	actualProducer, ok := actualResult.(*producer)
	assert.True(t, ok)
	assert.Equal(t, config.ProducerCount, actualProducer.producerCount)
	assert.Equal(t, config.Sender, actualProducer.sender)
	assert.Equal(t, config.Events, actualProducer.events)
	assert.Equal(t, config.WorkerPool, actualProducer.workerPool)
	assert.Equal(t, config.Repo, actualProducer.repo)
}

func TestProducer_Start(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockEventSender(ctrl)
	repo := mocks.NewMockEventRepo(ctrl)
	events := make(chan model.MessageEvent, 5)
	workerPool := workerpool.New(2)
	config := Config{
		2,
		sender,
		events,
		workerPool,
		repo,
	}
	producer := NewKafkaProducer(config)

	repo.EXPECT().Lock(gomock.Any()).AnyTimes()

	ctx, cancel := context.WithCancel(context.Background())
	producer.Start(ctx)
	defer producer.Close()
	cancel()
}

func TestProducer_SuccessfulSend(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	events := make(chan model.MessageEvent, 1)
	events <- testEvent

	workerPool := workerpool.New(1)
	config := Config{
		1,
		sender,
		events,
		workerPool,
		repo,
	}
	producer := NewKafkaProducer(config)

	sender.EXPECT().Send(gomock.Eq(&testEvent)).Return(nil).Times(1)
	repo.EXPECT().Remove(gomock.Any()).Times(1)
	ctx, cancel := context.WithCancel(context.Background())

	producer.Start(ctx)

	wait(events)
	defer producer.Close()
	cancel()
	require.Equal(t, 0, len(events), "Event was not sent")
}

func TestProducer_UnsuccessfulSent(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	events := make(chan model.MessageEvent, 1)
	events <- testEvent

	workerPool := workerpool.New(1)
	config := Config{
		1,
		sender,
		events,
		workerPool,
		repo,
	}
	producer := NewKafkaProducer(config)

	sender.EXPECT().Send(gomock.Eq(&testEvent)).Return(errors.New("Error during sending event")).Times(1)
	repo.EXPECT().Unlock(gomock.Any()).Times(1)

	ctx, cancel := context.WithCancel(context.Background())
	producer.Start(ctx)
	defer producer.Close()
	wait(events)
	cancel()
	require.Equal(t, 0, len(events), "Event was not sent")
}

func wait(events chan model.MessageEvent) {
	for len(events) > 0 {
		time.Sleep(50 * time.Millisecond)
	}
}