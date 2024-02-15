package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"transport-api/internal/adapters/grpccore"
	"transport-api/internal/app"
	"transport-api/internal/ports/httpserver"
	"transport-api/pkg/tokenizer"
)

const (
	dockerConfigFile = "config/config-docker.yml"
	localConfigFile  = "config/config-local.yml"
)

//	@title			BRM API
//	@version		1.0
//	@description	Swagger документация к API
//	@host			localhost:8090
//	@BasePath		/api/v1

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

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("reading config: %s", err.Error())
	}
	if err := godotenv.Load("config/.env"); err != nil {
		log.Fatalf("unable to load .env file: %s", err.Error())
	}

	coreClient, err := grpccore.NewCoreClient(ctx, fmt.Sprintf("%s:%d",
		viper.GetString("grpc-core-client.host"),
		viper.GetInt("grpc-core-client.port")))
	if err != nil {
		log.Fatal("create grpc core client: ", err.Error())
	}

	a := app.NewApp(coreClient)
	tkn := tokenizer.New(os.Getenv("SIGNKEY"))

	srv := httpserver.New(fmt.Sprintf("%s:%d",
		viper.GetString("http-server.host"),
		viper.GetInt("http-server.port")),
		a, tkn)

	go func() {
		if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("listening server: ", err.Error())
		}
	}()

	// preparing graceful shutdown
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGINT)

	// waiting for Ctrl+C
	<-osSignals

	shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second) // 30s timeout to finish all active connections
	defer cancel()

	_ = srv.Shutdown(shutdownCtx)
}
