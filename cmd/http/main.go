package main

import (
	http_handler "checkout-case/internal/adapter/handler/http"
	repository "checkout-case/internal/adapter/repository/mongodb"
	cart_service "checkout-case/internal/core/service/cart"
	promotion_service "checkout-case/internal/core/service/promotion"
	http_platform "checkout-case/internal/platform/server/http"
	"checkout-case/pkg/config"
	"checkout-case/pkg/logger"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("[FATAL] %s", err)
	}

	l := logger.RegisterLogger(config.Cfg.LoggerConfig).Sugar()

	l.Info("checkout-case is started...")
	l.Infof("Runtime config %+v", config.Cfg)
	l.Infof("Go runtime is %s", runtime.Version())

	var (
		err         error
		errRestChan = make(chan error, 10)
	)

	mongoCollection, err := repository.Connection()
	if err != nil {
		l.Error(fmt.Errorf("mongo connection error: %w", err))
		return
	}

	cartRepo := repository.NewCartRepository(mongoCollection)

	if err := cartRepo.Create(); err != nil {
		l.Error(fmt.Errorf("initil cart error: %w", err))
		return
	}
	l.Info("init cart successfully created")

	promotionService := promotion_service.NewPromotionService()
	cartService := cart_service.NewCartService(cartRepo, promotionService)

	restHandler := http_handler.NewRestHandler(cartService)

	restServer := http_platform.NewRestServer(restHandler)

	err = restServer.Start(errRestChan)
	if err != nil {
		l.Errorf("rest server init error: %v", err)
		return
	}
	l.Info("rest server successfully started")

	signalChan := make(chan os.Signal, 1)
	// syscall package and SIGTERM may be removed for Windows
	// os.Interrupt (SIGINT) and os.Kill (SIGKILL) are only signal values guaranteed to be present on all systems
	// But os.Kill cannot be trapped by a program
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case err := <-errRestChan:
			if err != nil {
				l.Errorf("Rest Server error: %+v", err)
			}
			l.Info("Shutdown completed...")

			os.Exit(0)
		case <-signalChan:
			l.Info("Signal received, shutting down...")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			restServer.GracefulShutdown(ctx)
			l.Info("rest server graceful shutdown is done")
		}
	}
}
