package main

import (
	file_handler "checkout-case/internal/adapter/handler/file"
	repository "checkout-case/internal/adapter/repository/mongodb"
	cart_service "checkout-case/internal/core/service/cart"
	promotion_service "checkout-case/internal/core/service/promotion"
	file_platform "checkout-case/internal/platform/server/file"
	"checkout-case/pkg/config"
	"checkout-case/pkg/logger"
	"fmt"
	"log"
	"runtime"
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
		errFileChan = make(chan error, 10)
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

	fileHandler := file_handler.NewFileHandler(cartService)

	fileServer := file_platform.NewFileServer(fileHandler)

	err = fileServer.Start(errFileChan)
	if err != nil {
		l.Errorf("file server init error: %v", err)
		return
	}
	l.Info("file server successfully started")
}
