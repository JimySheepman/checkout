//go:generate mockgen -destination=./mocks/cart_mock.go -source=cart.go

package port

import (
	"checkout-case/internal/core/domain"
	"checkout-case/internal/core/models"
	"context"
)

type CartService interface {
	AddItemToCart(ctx context.Context, req *models.AddItemServiceRequest) error
	AddVasItemToItem(ctx context.Context, req *models.AddVasItemToItemServiceRequest) error
	RemoveItemFromCart(ctx context.Context, itemId int) error
	ResetCart(ctx context.Context) error
	DisplayCart(ctx context.Context) (*models.DisplayCartServiceResponse, error)
}

type CartRepository interface {
	ResetCart() error
	GetCart() (*domain.Cart, error)
	AddItem(item *domain.Item) error
	RemoveItem(item *domain.Item) error
	FindItemByItemIdFromCart(itemId int) (*domain.Item, error)
	UpdateItemQuantity(item *domain.Item, req *models.AddItemServiceRequest) error
	UpdateVasItemQuantity(item *domain.Item, vasItem *domain.VasItem, req *models.AddVasItemToItemServiceRequest) error
	AddVasItemToItemByItemID(itemId string, vasItem *domain.VasItem) error
}
