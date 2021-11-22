package logging

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strconv"
)

//DetailedInfoRequestUnaryServerInterceptor interceptor for adding details about request
func DetailedInfoRequestUnaryServerInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		isDetailedHeader := md.Get("is-detailed-info")
		if (len(isDetailedHeader)) > 0 {
			isDetailed, _ := strconv.ParseBool(isDetailedHeader[0])
			if isDetailed {
				log.Info().Str("method", info.FullMethod).
					Interface("metadata", md).
					Interface("request body", req).
					Msg("Logged detailed request info")
			}
		}
	}

	h, err := handler(ctx, req)
	return h, err
}

//DetailedInfoResponseUnaryServerInterceptor interceptor for adding details about response
func DetailedInfoResponseUnaryServerInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	h, err := handler(ctx, req)

	if ok {
		isDetailedHeader := md.Get("is-detailed-info")
		if (len(isDetailedHeader)) > 0 {
			isDetailed, _ := strconv.ParseBool(isDetailedHeader[0])
			if isDetailed {
				if err != nil {
					log.Error().Str("method", info.FullMethod).
						Err(err).
						Msg("Request failed")
				} else {
					log.Info().Str("method", info.FullMethod).
						Interface("response", h).
						Msg("Request finished successfully")
				}
			}
		}
	}

	return h, err
}
