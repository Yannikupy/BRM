package grpcserver

import (
	"auth/internal/app"
	"auth/internal/ports/grpcserver/pb"
	"auth/pkg/logger"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

type Server struct {
	App app.App
	pb.AuthServiceServer
}

func New(a app.App, logs logger.Logger) *grpc.Server {
	chain := grpcmiddleware.ChainUnaryServer(
		panicInterceptor(logs),
		loggerInterceptor(logs))

	s := grpc.NewServer(grpc.UnaryInterceptor(chain))
	pb.RegisterAuthServiceServer(s, &Server{
		App:               a,
		AuthServiceServer: nil,
	})
	return s
}
