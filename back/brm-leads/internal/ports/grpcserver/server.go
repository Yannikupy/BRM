package grpcserver

import (
	"brm-leads/internal/app"
	"brm-leads/internal/ports/grpcserver/pb"
	"brm-leads/pkg/logger"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

type Server struct {
	App app.App
	pb.LeadsServiceServer
}

func New(a app.App, logs logger.Logger) *grpc.Server {
	chain := grpcmiddleware.ChainUnaryServer(
		panicInterceptor(logs),
		loggerInterceptor(logs))

	s := grpc.NewServer(grpc.UnaryInterceptor(chain))
	pb.RegisterLeadsServiceServer(s, &Server{
		App:                a,
		LeadsServiceServer: nil,
	})
	return s
}
