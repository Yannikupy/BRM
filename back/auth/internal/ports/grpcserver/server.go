package grpcserver

import (
	"auth/internal/app"
	"auth/internal/ports/grpcserver/pb"
	"google.golang.org/grpc"
)

type Server struct {
	App app.App
	pb.AuthServiceServer
}

func New(a app.App) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &Server{
		App:               a,
		AuthServiceServer: nil,
	})
	return s
}
