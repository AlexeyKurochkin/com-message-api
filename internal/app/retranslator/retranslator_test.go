package retranslator

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/com-message-api/internal/mocks"
	"testing"
	"time"
)

func TestRetranslator_Start(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	repo.EXPECT().Lock(gomock.Any()).AnyTimes()

	cfg := Config{
		ConsumerCount:  2,
		ConsumeSize:    10,
		ConsumeTimeout: time.Second,
		ProducerCount:  2,
		WorkerCount:    2,
		ChannelSize:    512,
		Repo:           repo,
		Sender:         sender,
	}
	retranslator := NewRetranslator(cfg)
	ctx, cancel := context.WithCancel(context.Background())

	retranslator.Start(ctx)
	cancel()
	retranslator.Close()
}
