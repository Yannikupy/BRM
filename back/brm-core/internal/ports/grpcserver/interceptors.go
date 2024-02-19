package grpcserver

import (
	"brm-core/pkg/logger"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func loggerInterceptor(logs logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			return nil, mapErrors(err)
		}
		logs.Info(logger.Fields{
			"Method": info.FullMethod,
		}, "got request")
		return resp, nil
	}
}

func panicInterceptor(logs logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		defer func() {
			if r := recover(); r != nil {
				logs.Error(logger.Fields{
					"Method": info.FullMethod,
				}, fmt.Sprintf("panic: %v", r))
			}
		}()
		resp, err = handler(ctx, req)
		if err != nil {
			return nil, mapErrors(err)
		}
		return resp, nil
	}
}
