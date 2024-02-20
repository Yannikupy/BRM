package grpcserver

import (
	"brm-core/internal/app"
	"brm-core/internal/ports/grpcserver/pb"
	"brm-core/pkg/logger"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

type Server struct {
	App app.App
	pb.CoreServiceServer
}

func New(a app.App, logs logger.Logger) *grpc.Server {
	chain := grpcmiddleware.ChainUnaryServer(
		panicInterceptor(logs),
		loggerInterceptor(logs))

	s := grpc.NewServer(grpc.UnaryInterceptor(chain))
	pb.RegisterCoreServiceServer(s, &Server{
		App:               a,
		CoreServiceServer: nil,
	})
	return s
}
