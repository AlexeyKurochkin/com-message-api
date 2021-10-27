package main

import (
	"github.com/ozonmp/com-message-api/internal/app/retranslator"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	cfg := retranslator.Config{
		ConsumerCount:  2,
		ConsumeSize:    10,
		ConsumeTimeout: 10 * time.Second,
		ProducerCount:  28,
		WorkerCount:    2,
		ChannelSize:    512,
	}

	ctx, cancel := context.WithCancel(context.Background())
	retranslator := retranslator.NewRetranslator(cfg)
	retranslator.Start(ctx)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	cancel()
	retranslator.Close()
}
