package logging

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func LogLevelUnaryServerInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		levels := md.Get("log-level")
		if (len(levels)) > 0 {
			log.Info().Msgf("Got log levels: %v", levels)
			if parsedLevel, err := zerolog.ParseLevel(levels[0]); err != nil {
				log.Error().Msg("Unable to parse log level from headers")
			} else {
				logger := log.Level(parsedLevel)
				ctx = logger.WithContext(ctx)
				log.Info().Msgf("Successfully applied log level '%v' from headers", parsedLevel)
			}
		}
	}
	h, err := handler(ctx, req)

	return h, err
}
