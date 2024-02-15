package main

import (
	"auth/cmd/server/factory"
	"auth/internal/app"
	"auth/internal/app/tokenizer"
	"auth/internal/ports/grpcserver"
	"auth/internal/ports/httpserver"
	"auth/internal/repo/authrepo"
	"auth/internal/repo/passrepo"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//	@title			BRM API
//	@version		1.0
//	@description	Swagger документация к API авторизации
//	@host			localhost:8082
//	@BasePath		/api/v1/auth

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

	passConn, err := factory.ConnectToPostgres(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() {
		if passConn != nil {
			_ = passConn.Close(ctx)
		}
	}()

	redisClient := factory.ConnectToRedis()
	if redisClient == nil {
		log.Fatal("unable to connect to redis")
	}
	defer func() {
		if redisClient != nil {
			_ = redisClient.Close()
		}
	}()

	authRepo, err := authrepo.New(
		redisClient,
		time.Duration(viper.GetInt("app.refresh-token-expiration")))
	if err != nil {
		log.Fatal(err.Error())
	}

	passRepo := passrepo.New(passConn)

	tkn := tokenizer.New(
		time.Duration(viper.GetInt("app.access-token-expiration")),
		[]byte(os.Getenv("SIGNKEY")),
	)

	a := app.New(
		os.Getenv("PASSWORDSALT"),
		viper.GetInt("app.refresh-token-length"),
		authRepo,
		passRepo,
		tkn,
	)

	grpcsrv := grpcserver.New(a)
	lis, err := factory.PrepareListener()
	if err != nil {
		log.Fatal(err.Error())
	}

	go func() {
		if err = grpcsrv.Serve(lis); err != nil {
			log.Fatal("starting grpc server: ", err.Error())
		}
	}()

	httpsrv := httpserver.New(fmt.Sprintf("%s:%d",
		viper.GetString("http-server.host"),
		viper.GetInt("http-server.port")),
		a)

	go func() {
		if err = httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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

	grpcsrv.GracefulStop()
	_ = httpsrv.Shutdown(shutdownCtx)
}
