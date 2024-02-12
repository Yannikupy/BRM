package main

import (
	"context"
	"errors"
	"fmt"
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
)

//	@title			BRM API
//	@version		1.0
//	@description	Swagger документация к API
//	@host			localhost:8080
//	@BasePath		/api/v1

func main() {
	ctx := context.Background()

	viper.SetConfigFile("config/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("reading config: %s", err.Error())
	}

	coreClient, err := grpccore.NewCoreClient(ctx, fmt.Sprintf("%s:%d",
		viper.GetString("grpc-core-client.host"),
		viper.GetInt("grpc-core-client.port")))
	if err != nil {
		log.Fatal("create grpc core client: ", err.Error())
	}

	a := app.NewApp(coreClient)

	srv := httpserver.New(fmt.Sprintf("%s:%d",
		viper.GetString("http-server.host"),
		viper.GetInt("http-server.port")),
		a)

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
