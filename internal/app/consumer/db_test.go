package consumer

import (
	"github.com/ozonmp/com-message-api/internal/mocks"
	"github.com/ozonmp/com-message-api/internal/model"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testEvent = model.MessageEvent{
	ID:     1,
	Type:   model.Created,
	Status: model.Deferred,
	Entity: &model.Message{},
}

func TestConsumer_NewDbConsumer(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	events := make(chan model.MessageEvent, 1)
	config := Config{
		1,
		2,
		10 * time.Second,
		events,
		repo,
	}

	actualResult := NewDbConsumer(config)
	actualConsumer, ok := actualResult.(*consumer)
	assert.True(t, ok)
	assert.Equal(t, config.ConsumerCount, actualConsumer.consumerCount)
	assert.Equal(t, config.BatchSize, actualConsumer.batchSize)
	assert.Equal(t, config.Timeout, actualConsumer.timeout)
	assert.Equal(t, config.Events, actualConsumer.events)
	assert.Equal(t, config.Repo, actualConsumer.repo)
}

func TestConsumer_Start(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	events := make(chan model.MessageEvent, 1)
	//events <- testEvent
	config := Config{
		1,
		2,
		1 * time.Second,
		events,
		repo,
	}
	repo.EXPECT().Lock(gomock.Any()).Return([]model.MessageEvent{testEvent}, nil).AnyTimes()
	ctx, cancel := context.WithCancel(context.Background())

	actualResult := NewDbConsumer(config)
	actualResult.Start(ctx)

	wait(events, config.Timeout)
	cancel()
	assert.Equal(t, 1, len(events), "Event was not consumed")
}

func wait(events chan model.MessageEvent, consumingTime time.Duration) {
	for i := 0; i < 3; i++ {
		time.Sleep(consumingTime)
		if len(events) > 0 {
			break
		}
	}
}
