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
	"transport-api/internal/adapters/grpcads"
	"transport-api/internal/adapters/grpccore"
	"transport-api/internal/adapters/grpcleads"
	"transport-api/internal/app"
	"transport-api/internal/ports/httpserver"
	"transport-api/pkg/logger"
	"transport-api/pkg/tokenizer"
)

const (
	dockerConfigFile = "config/config-docker.yml"
	localConfigFile  = "config/config-local.yml"
)

//	@title						BRM API
//	@version					1.0
//	@description				Swagger документация к API
//	@host						localhost:8090
//	@BasePath					/api/v1
//	@SecurityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	ctx := context.Background()
	logs := logger.New()

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
		logs.Fatal(nil, fmt.Sprintf("reading config: %s", err.Error()))
	}
	if err := godotenv.Load("config/.env"); err != nil {
		log.Fatalf("unable to load .env file: %s", err.Error())
	}

	coreClient, err := grpccore.NewCoreClient(ctx, fmt.Sprintf("%s:%d",
		viper.GetString("grpc-core-client.host"),
		viper.GetInt("grpc-core-client.port")))
	if err != nil {
		logs.Fatal(nil, fmt.Sprintf("create grpc core client: %s", err.Error()))
	}

	adsClient, err := grpcads.NewAdsClient(ctx, fmt.Sprintf("%s:%d",
		viper.GetString("grpc-ads-client.host"),
		viper.GetInt("grpc-ads-client.port"),
	))
	if err != nil {
		logs.Fatal(nil, fmt.Sprintf("create ads core client: %s", err.Error()))
	}

	leadsClient, err := grpcleads.NewLeadsClient(ctx, fmt.Sprintf("%s:%d",
		viper.GetString("grpc-leads-client.host"),
		viper.GetInt("grpc-leads-client.port"),
	))
	if err != nil {
		logs.Fatal(nil, fmt.Sprintf("create leads core client: %s", err.Error()))
	}

	a := app.NewApp(coreClient, adsClient, leadsClient)
	tkn := tokenizer.New(os.Getenv("SIGNKEY"))

	srv := httpserver.New(
		fmt.Sprintf("%s:%d",
			viper.GetString("http-server.host"),
			viper.GetInt("http-server.port")),
		fmt.Sprintf("http://%s:%d",
			viper.GetString("origins.web.host"),
			viper.GetInt("origins.web.port")),
		a, tkn, logs)

	go func() {
		if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logs.Fatal(nil, fmt.Sprintf("listening server: %s", err.Error()))
		}
	}()

	logs.Info(nil, "service transport-api successfully started")

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
