package file

import (
	"checkout-case/internal/core/models"
	"checkout-case/internal/core/port"
	"checkout-case/pkg/logger"
	"context"
	"encoding/json"
	"fmt"
)

type fileHandler struct {
	cartService port.CartService
}

func NewFileHandler(cartService port.CartService) *fileHandler {
	return &fileHandler{
		cartService: cartService,
	}
}

func (h *fileHandler) AddItemHandler(ctx context.Context, input string) (string, error) {
	l := logger.FromCtx(ctx).Sugar()

	item := &models.AddItemCommandRequest{}
	if err := json.Unmarshal([]byte(input), item); err != nil {
		l.Error(err)
		return "", err
	}

	if err := h.cartService.AddItemToCart(ctx, populateAddItemCommandRequestToAddItemServiceRequest(item)); err != nil {
		l.Errorf("add item to cart service error: %v", err)
		return "", err
	}

	return "item was added to cart successfully", nil
}

func populateAddItemCommandRequestToAddItemServiceRequest(item *models.AddItemCommandRequest) *models.AddItemServiceRequest {
	return &models.AddItemServiceRequest{
		ItemID:     item.ItemID,
		CategoryID: item.CategoryID,
		SellerID:   item.SellerID,
		Price:      item.Price,
		Quantity:   item.Quantity,
	}
}

func (h *fileHandler) AddVasItemToItemHandler(ctx context.Context, input string) (string, error) {
	l := logger.FromCtx(ctx).Sugar()

	vasI := &models.AddVasItemToItemCommandRequest{}
	if err := json.Unmarshal([]byte(input), vasI); err != nil {
		l.Error(err)
		return "", err
	}

	if err := h.cartService.AddVasItemToItem(ctx, populateAddVasItemToItemCommandRequestToAddVasItemToItemServiceRequest(vasI)); err != nil {
		l.Errorf("add vasItem to item service service error: %v", err)
		return "", err
	}

	return "vasItem was added to item successfully", nil
}

func populateAddVasItemToItemCommandRequestToAddVasItemToItemServiceRequest(vasI *models.AddVasItemToItemCommandRequest) *models.AddVasItemToItemServiceRequest {
	return &models.AddVasItemToItemServiceRequest{
		ItemID:     vasI.ItemID,
		VasItemID:  vasI.VasItemID,
		CategoryID: vasI.CategoryID,
		SellerID:   vasI.SellerID,
		Price:      vasI.Price,
		Quantity:   vasI.Quantity,
	}
}

func (h *fileHandler) RemoveItemHandler(ctx context.Context, input string) (string, error) {
	l := logger.FromCtx(ctx).Sugar()

	item := &models.RemoveItemCommandRequest{}
	if err := json.Unmarshal([]byte(input), item); err != nil {
		return "", fmt.Errorf("remove item command request Unmarshal error: %w", err)
	}

	if err := h.cartService.RemoveItemFromCart(ctx, item.ItemID); err != nil {
		l.Errorf("remove item from cart service error: %v", err)
		return "", err
	}

	return "item was removed to cart successfully", nil
}

func (h *fileHandler) ResetCartHandler(ctx context.Context, input string) (string, error) {
	l := logger.FromCtx(ctx).Sugar()

	if err := h.cartService.ResetCart(ctx); err != nil {
		l.Errorf("reset cart service error: %v", err)
		return "", err
	}

	return "cart was reset successfully", nil
}

func (h *fileHandler) DisplayCartHandler(ctx context.Context, input string) (string, error) {
	l := logger.FromCtx(ctx).Sugar()

	resp, err := h.cartService.DisplayCart(ctx)
	if err != nil {
		l.Errorf("display cart service error: %v", err)
		return "", err
	}

	return resp.ToString(), nil
}
