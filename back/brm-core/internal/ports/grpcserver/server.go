package grpcserver

import (
	"brm-core/internal/app"
	"brm-core/internal/ports/grpcserver/pb"
	"google.golang.org/grpc"
)

type Server struct {
	App app.App
	pb.CoreServiceServer
}

func New(a app.App) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterCoreServiceServer(s, &Server{
		App:               a,
		CoreServiceServer: nil,
	})
	return s
}
