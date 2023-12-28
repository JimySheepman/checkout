package http

import (
	"checkout-case/internal/core/models"
	"checkout-case/internal/core/port"
	"checkout-case/pkg/logger"
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

const (
	loggerRequestIdKey  = "requestId"
	appIsHealthyMessage = "program is work"

	RemoveItemPathParam = "itemId"
)

type restHandler struct {
	cartService port.CartService
}

func NewRestHandler(cartService port.CartService) *restHandler {
	return &restHandler{
		cartService: cartService,
	}
}

func (h *restHandler) HealthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, appIsHealthyMessage)
}

func (h *restHandler) AddItemHandler(c echo.Context) error {
	l := logger.GetLogger().With(
		zap.String(loggerRequestIdKey, c.Response().Header().Get(echo.HeaderXRequestID)),
	)
	ctx := logger.WithCtx(context.TODO(), l)

	item := &models.AddItemRestRequest{}
	if err := json.NewDecoder(c.Request().Body).Decode(item); err != nil {
		l.Sugar().Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.cartService.AddItemToCart(ctx, populateAddItemRestRequestToAddItemServiceRequest(item)); err != nil {
		l.Sugar().Errorf("add item to cart service error: %v", err)
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusCreated)
}

func populateAddItemRestRequestToAddItemServiceRequest(item *models.AddItemRestRequest) *models.AddItemServiceRequest {
	return &models.AddItemServiceRequest{
		ItemID:     item.ItemID,
		CategoryID: item.CategoryID,
		SellerID:   item.SellerID,
		Price:      item.Price,
		Quantity:   item.Quantity,
	}
}

func (h *restHandler) AddVasItemToItemHandler(c echo.Context) error {
	l := logger.GetLogger().With(
		zap.String(loggerRequestIdKey, c.Response().Header().Get(echo.HeaderXRequestID)),
	)
	ctx := logger.WithCtx(context.TODO(), l)

	vasI := &models.AddVasItemToItemRestRequest{}
	if err := json.NewDecoder(c.Request().Body).Decode(vasI); err != nil {
		l.Sugar().Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.cartService.AddVasItemToItem(ctx, populateAddVasItemToItemRestRequestToAddVasItemToItemServiceRequest(vasI)); err != nil {
		l.Sugar().Errorf("add vasItem to item service service error: %v", err)
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusCreated)
}

func populateAddVasItemToItemRestRequestToAddVasItemToItemServiceRequest(vasI *models.AddVasItemToItemRestRequest) *models.AddVasItemToItemServiceRequest {
	return &models.AddVasItemToItemServiceRequest{
		ItemID:     vasI.ItemID,
		VasItemID:  vasI.VasItemID,
		CategoryID: vasI.CategoryID,
		SellerID:   vasI.SellerID,
		Price:      vasI.Price,
		Quantity:   vasI.Quantity,
	}
}

func (h *restHandler) RemoveItemHandler(c echo.Context) error {
	l := logger.GetLogger().With(
		zap.String(loggerRequestIdKey, c.Response().Header().Get(echo.HeaderXRequestID)),
	)
	ctx := logger.WithCtx(context.TODO(), l)

	itemId, err := strconv.Atoi(c.Param(RemoveItemPathParam))
	if err != nil {
		l.Sugar().Error(err)
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.cartService.RemoveItemFromCart(ctx, itemId); err != nil {
		l.Sugar().Errorf("remove item from cart service error: %v", err)
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}

func (h *restHandler) ResetCartHandler(c echo.Context) error {
	l := logger.GetLogger().With(
		zap.String(loggerRequestIdKey, c.Response().Header().Get(echo.HeaderXRequestID)),
	)
	ctx := logger.WithCtx(context.TODO(), l)

	if err := h.cartService.ResetCart(ctx); err != nil {
		l.Sugar().Errorf("reset cart service error: %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func (h *restHandler) DisplayCartHandler(c echo.Context) error {
	l := logger.GetLogger().With(
		zap.String(loggerRequestIdKey, c.Response().Header().Get(echo.HeaderXRequestID)),
	)
	ctx := logger.WithCtx(context.TODO(), l)

	resp, err := h.cartService.DisplayCart(ctx)
	if err != nil {
		l.Sugar().Errorf("display cart service error: %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, resp)
}
