package main

import (
	"brm-core/cmd/server/factory"
	"brm-core/internal/adapters/grpcauth"
	"brm-core/internal/app"
	"brm-core/internal/ports/grpcserver"
	"brm-core/internal/repo"
	"context"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	dockerConfigFile = "config/config-docker.yml"
	localConfigFile  = "config/config-local.yml"
)

func main() {
	ctx := context.Background()

	isDocker := flag.Bool("docker", false, "flag if this project is running in docker container")
	flag.Parse()
	var configPath string
	if *isDocker {
		configPath = dockerConfigFile
	} else {
		configPath = localConfigFile
	}

	if err := factory.SetConfigs(configPath); err != nil {
		log.Fatal(err.Error())
	}

	coreRepo, err := factory.ConnectToPostgres(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	authClient, err := grpcauth.NewAuthClient(ctx, fmt.Sprintf("%s:%d",
		viper.GetString("grpc-auth-client.host"),
		viper.GetInt("grpc-auth-client.port")))

	a := app.New(repo.New(coreRepo), authClient)

	srv := grpcserver.New(a)
	lis, err := factory.PrepareListener()
	if err != nil {
		log.Fatal(err.Error())
	}

	// preparing graceful shutdown
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGINT)

	go func() {
		if err = srv.Serve(lis); err != nil {
			log.Fatal("starting grpc server: ", err.Error())
		}
	}()

	<-osSignals
	srv.GracefulStop()
}
