//go:build integration
// +build integration

package mongodb

import (
	"checkout-case/internal/core/domain"
	"fmt"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func setupTestRepository() *cartRepository {
	return &cartRepository{
		client:     db,
		collection: db.Database("shopping").Collection("carts"),
	}
}

func TestCartRepository_Create(t *testing.T) {
	repo := setupTestRepository()

	err := repo.Create()
	require.Nil(t, err)

	_, err = repo.GetCart()
	require.Nil(t, err)
}

func TestCartRepository_GetCart(t *testing.T) {
	repo := setupTestRepository()

	err := repo.Create()
	require.Nil(t, err)

	_, err = repo.GetCart()
	require.Nil(t, err)
}

func TestCartRepository_AddItem(t *testing.T) {
	now := time.Now()
	repo := setupTestRepository()

	err := repo.Create()
	require.Nil(t, err)

	item := &domain.Item{
		ID:         primitive.NewObjectID().String(),
		ItemId:     1,
		CategoryId: 1,
		SellerId:   1,
		CreatedAt:  now,
		UpdatedAt:  now,
		Type:       0,
		Price:      1,
		Quantity:   1,
		VasItems:   make([]*domain.VasItem, 0),
	}

	err = repo.AddItem(item)
	require.Nil(t, err)

	actual, err := repo.FindItemByItemIdFromCart(item.ItemId)
	require.Nil(t, err)

	fmt.Println(actual)
}

func TestCartRepository_UpdateItemQuantity(t *testing.T) {
	now := time.Now()
	repo := setupTestRepository()

	err := repo.Create()
	require.Nil(t, err)

	item := &domain.Item{
		ID:         primitive.NewObjectID().String(),
		ItemId:     1,
		CategoryId: 1,
		SellerId:   1,
		CreatedAt:  now,
		UpdatedAt:  now,
		Type:       0,
		Price:      1,
		Quantity:   1,
		VasItems:   make([]*domain.VasItem, 0),
	}

	expectedQuantity := item.Quantity

	err = repo.AddItem(item)
	require.Nil(t, err)

	item.Quantity += 2

	err = repo.UpdateItemQuantity(item, nil)

	actual, err := repo.FindItemByItemIdFromCart(item.ItemId)
	require.Nil(t, err)
	require.NotEqual(t, expectedQuantity, actual.Quantity)
}

func TestCartRepository_AddVasItemToItemByItemID(t *testing.T) {
	now := time.Now()

	repo := setupTestRepository()

	err := repo.Create()
	require.Nil(t, err)

	item := &domain.Item{
		ID:         primitive.NewObjectID().String(),
		ItemId:     1,
		CategoryId: 1,
		SellerId:   1,
		CreatedAt:  now,
		UpdatedAt:  now,
		Type:       0,
		Price:      1,
		Quantity:   1,
		VasItems:   make([]*domain.VasItem, 0),
	}

	err = repo.AddItem(item)
	require.Nil(t, err)

	vasItem := &domain.VasItem{
		ID:         primitive.NewObjectID().String(),
		VasItemId:  1,
		CategoryId: 1,
		SellerId:   1,
		CreatedAt:  now,
		UpdatedAt:  now,
		Price:      1,
		Quantity:   1,
	}

	err = repo.AddVasItemToItemByItemID(item.ID, vasItem)
	require.Nil(t, err)

	actual, err := repo.FindVasItemByVasItemIdFromItem(item.ItemId)
	require.Nil(t, err)
	require.Equal(t, vasItem.VasItemId, actual.VasItemId)
}

func TestCartRepository_UpdateVasItemQuantity(t *testing.T) {
	now := time.Now()

	repo := setupTestRepository()

	err := repo.Create()
	require.Nil(t, err)

	item := &domain.Item{
		ID:         primitive.NewObjectID(),
		ItemId:     1,
		CategoryId: 1,
		SellerId:   1,
		CreatedAt:  now,
		UpdatedAt:  now,
		Type:       0,
		Price:      1,
		Quantity:   1,
		VasItems:   make([]*domain.VasItem, 0),
	}

	err = repo.AddItem(item)
	require.Nil(t, err)

	vasItem := &domain.VasItem{
		ID:         primitive.NewObjectID().String(),
		VasItemId:  1,
		CategoryId: 1,
		SellerId:   1,
		CreatedAt:  now,
		UpdatedAt:  now,
		Price:      1,
		Quantity:   1,
	}

	err = repo.AddVasItemToItemByItemID(item.ID, vasItem)
	require.Nil(t, err)

	vasItem.Quantity += 2

	err = repo.UpdateVasItemQuantity(item, vasItem, nil)
	require.Nil(t, err)

	actual, err := repo.FindVasItemByVasItemIdFromItem(item.ItemId)
	require.Nil(t, err)
	require.Equal(t, vasItem.Quantity, actual.Quantity)
}

func TestCartRepository_ResetCart(t *testing.T) {
	now := time.Now()
	repo := setupTestRepository()

	err := repo.Create()
	require.Nil(t, err)

	item := &domain.Item{
		ID:         primitive.NewObjectID().String(),
		ItemId:     1,
		CategoryId: 1,
		SellerId:   1,
		CreatedAt:  now,
		UpdatedAt:  now,
		Type:       0,
		Price:      1,
		Quantity:   1,
		VasItems:   make([]*domain.VasItem, 0),
	}

	err = repo.AddItem(item)
	require.Nil(t, err)

	err = repo.ResetCart()
	require.Nil(t, err)

	actual, err := repo.FindItemByItemIdFromCart(item.ItemId)
	require.Error(t, err)
	require.Nil(t, actual)
}

func TestCartRepository_RemoveItem(t *testing.T) {
	now := time.Now()
	repo := setupTestRepository()

	err := repo.Create()
	require.Nil(t, err)

	item := &domain.Item{
		ID:         primitive.NewObjectID().String(),
		ItemId:     1,
		CategoryId: 1,
		SellerId:   1,
		CreatedAt:  now,
		UpdatedAt:  now,
		Type:       0,
		Price:      1,
		Quantity:   1,
		VasItems:   make([]*domain.VasItem, 0),
	}

	err = repo.AddItem(item)
	require.Nil(t, err)

	err = repo.RemoveItem(item)
	require.Nil(t, err)

	actual, err := repo.FindItemByItemIdFromCart(item.ItemId)
	require.Error(t, err)
	require.Nil(t, actual)
}

func TestCartRepository_FindItemByItemIdFromCart(t *testing.T) {
	now := time.Now()
	repo := setupTestRepository()

	err := repo.Create()
	require.Nil(t, err)

	item := &domain.Item{
		ID:         primitive.NewObjectID().String(),
		ItemId:     1,
		CategoryId: 1,
		SellerId:   1,
		CreatedAt:  now,
		UpdatedAt:  now,
		Type:       0,
		Price:      1,
		Quantity:   1,
		VasItems:   make([]*domain.VasItem, 0),
	}

	err = repo.AddItem(item)
	require.Nil(t, err)

	actual, err := repo.FindItemByItemIdFromCart(item.ItemId)
	require.Nil(t, err)

	fmt.Println(actual)
}

func TestCartRepository_FindVasItemByVasItemIdFromItem(t *testing.T) {
	now := time.Now()

	repo := setupTestRepository()

	err := repo.Create()
	require.Nil(t, err)

	item := &domain.Item{
		ID:         primitive.NewObjectID().String(),
		ItemId:     1,
		CategoryId: 1,
		SellerId:   1,
		CreatedAt:  now,
		UpdatedAt:  now,
		Type:       0,
		Price:      1,
		Quantity:   1,
		VasItems:   make([]*domain.VasItem, 0),
	}

	err = repo.AddItem(item)
	require.Nil(t, err)

	vasItem := &domain.VasItem{
		ID:         primitive.NewObjectID().String(),
		VasItemId:  1,
		CategoryId: 1,
		SellerId:   1,
		CreatedAt:  now,
		UpdatedAt:  now,
		Price:      1,
		Quantity:   1,
	}

	err = repo.AddVasItemToItemByItemID(item.ID, vasItem)
	require.Nil(t, err)

	actual, err := repo.FindVasItemByVasItemIdFromItem(item.ItemId)
	require.Nil(t, err)
	require.Equal(t, vasItem.VasItemId, actual.VasItemId)
}
