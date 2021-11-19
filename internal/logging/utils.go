package logging

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//GetLoggerAssociatedWithCtx returns logger associated with context with module and name labels
func GetLoggerAssociatedWithCtx(ctx context.Context, moduleName string, methodName string) *zerolog.Logger {
	logger := log.Ctx(ctx).With().Str("module", moduleName).Str("method", methodName).Logger()
	return &logger
}
