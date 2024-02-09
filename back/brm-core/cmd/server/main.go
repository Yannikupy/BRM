package main

import (
	"brm-core/cmd/server/factory"
	"brm-core/internal/app"
	"brm-core/internal/ports/grpcserver"
	"brm-core/internal/repo"
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ctx := context.Background()
	if err := factory.SetConfigs(); err != nil {
		log.Fatal(err.Error())
	}

	coreRepo, err := factory.ConnectToPostgres(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	a := app.New(repo.New(coreRepo))

	shards, err := factory.ConnectToRabbitmq()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() {
		for _, s := range shards {
			s.Close()
		}
	}()

	// preparing graceful shutdown
	sigQuit := make(chan os.Signal, 1)
	defer close(sigQuit)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)
	shutdownChans := make([]chan struct{}, len(shards))
	for i := range shutdownChans {
		shutdownChans[i] = make(chan struct{})
	}

	wg := new(sync.WaitGroup)
	for i := range shards {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			shards[i].HandleJobs(ctx, a, shutdownChans[i])
		}(i)
	}

	srv := grpcserver.New(a)
	lis, err := factory.PrepareListener()
	if err != nil {
		log.Fatal(err.Error())
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		_ = srv.Serve(lis)
	}()

	select {
	case <-sigQuit:
		srv.GracefulStop()
		for i := range shutdownChans {
			shutdownChans[i] <- struct{}{}
		}
	}

	wg.Wait()
}
