//go:generate mockgen -destination=../../../mocks/handler_mock.go -source=handler.go

package port

import (
	"context"
	"github.com/labstack/echo/v4"
)

type RestHandlerClient interface {
	HealthCheckHandler(c echo.Context) error
	AddItemHandler(c echo.Context) error
	AddVasItemToItemHandler(c echo.Context) error
	RemoveItemHandler(c echo.Context) error
	ResetCartHandler(c echo.Context) error
	DisplayCartHandler(c echo.Context) error
}

type FileHandlerClient interface {
	AddItemHandler(ctx context.Context, input string) (string, error)
	AddVasItemToItemHandler(ctx context.Context, input string) (string, error)
	RemoveItemHandler(ctx context.Context, input string) (string, error)
	ResetCartHandler(ctx context.Context, input string) (string, error)
	DisplayCartHandler(ctx context.Context, input string) (string, error)
}
