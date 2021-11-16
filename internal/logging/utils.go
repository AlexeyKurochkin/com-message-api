package logging

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func CreateLogger(ctx context.Context, moduleName string, methodName string) *zerolog.Logger {
	logger := log.Ctx(ctx).With().Str("module", moduleName).Str("method", methodName).Logger()
	return &logger
}
