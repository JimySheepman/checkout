package http

import (
	"checkout-case/internal/core/port"
	"checkout-case/pkg/config"
	"context"
	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const pprofEnabled = 1

type RestServer struct {
	e                 *echo.Echo
	restHandlerClient port.RestHandlerClient
}

func NewRestServer(restHandlerClient port.RestHandlerClient) *RestServer {
	return &RestServer{
		e:                 echo.New(),
		restHandlerClient: restHandlerClient,
	}
}

func (s *RestServer) Start(errChan chan error) error {
	_, cancel := context.WithCancel(context.Background())

	s.setMiddlewares()
	s.setRoutes()

	go func() {
		defer cancel()
		errChan <- s.e.Start(config.Cfg.Server.RestServer.Addr)
	}()

	return nil
}

func (s *RestServer) setMiddlewares() {
	s.e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.RequestID(),
	)
}

func (s *RestServer) setRoutes() {
	s.e.GET("/healthcheck", s.restHandlerClient.HealthCheckHandler)

	if config.Cfg.Server.RestServer.PprofEnable == pprofEnabled {
		pprof.Register(s.e, "/pprof")
	}

	prefixRouter := s.e.Group("/api/v1")
	{
		prefixRouter.POST("/item", s.restHandlerClient.AddItemHandler)
		prefixRouter.POST("/item/vas", s.restHandlerClient.AddVasItemToItemHandler)
		prefixRouter.DELETE("/item/:itemId", s.restHandlerClient.RemoveItemHandler)
		prefixRouter.DELETE("/cart", s.restHandlerClient.ResetCartHandler)
		prefixRouter.GET("/cart", s.restHandlerClient.DisplayCartHandler)
	}
}

func (s *RestServer) GracefulShutdown(ctx context.Context) {
	if err := s.e.Shutdown(ctx); err != nil {
		s.e.Logger.Fatal(err)
	}
}
