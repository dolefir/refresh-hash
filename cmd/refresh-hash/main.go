package main

import (
	"context"
	"errors"
	"flag"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/dolefir/refresh-hash/config"
	gen "github.com/dolefir/refresh-hash/gen/proto"
	"github.com/dolefir/refresh-hash/logger"
	inmemRepository "github.com/dolefir/refresh-hash/repository/inmem"
	"github.com/dolefir/refresh-hash/server/grpc/handler"
	"github.com/dolefir/refresh-hash/server/restapi"
	hashesHandler "github.com/dolefir/refresh-hash/server/restapi/handlers"
	hashService "github.com/dolefir/refresh-hash/services/hashes"
	"github.com/dolefir/refresh-hash/task"
	"google.golang.org/grpc"
)

func main() {
	cfgPath := flag.String("config", "config.yaml", "Configuration file")
	flag.Parse()
	cfg := config.NewConfig(*cfgPath)

	log := logger.NewLogger((*logger.CFGLogger)(&cfg.Logger), nil)

	hashRepo := inmemRepository.NewRepository()
	hashSrv := hashService.NewService(hashRepo, log)
	hashHdl := hashesHandler.NewHandler(hashSrv)

	ticker := task.NewRefreshTicker(cfg.Ticker.Timer, cfg.Ticker.Timeout, hashSrv, log)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// gRPC setup.
	list, err := net.Listen(cfg.APIServer.HTTP.Network, cfg.APIServer.HTTP.AddrGrpc)
	if err != nil {
		log.Fatal(err)
	}

	grpcHandler := handler.NewHashService(hashSrv)
	serviceRegistrar := grpc.NewServer()
	gen.RegisterHashServiceServer(serviceRegistrar, grpcHandler)
	go func() {
		if err := serviceRegistrar.Serve(list); err != nil {
			log.Fatal(err)
		}
	}()

	// REST API setup.
	api := restapi.NewAPI(hashHdl, cfg.APIServer, log)

	go func() {
		if err := api.ListenAndServe(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error(err)
		}
	}()

	// Ticker setup.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := ticker.Start(ctx); err != nil {
			log.Error(err)
		}
		wg.Done()
	}()

	// Metrics can also be configured
	// in this selection...

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	log.Info("shutdown...")
	cancel()

	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := api.Shutdown(shutdownCtx); err != nil {
		log.Fatal("server forced to shutdown: ", err)
	}

	serviceRegistrar.GracefulStop()

	wg.Wait()

	log.Info("successfully stopped")
}
