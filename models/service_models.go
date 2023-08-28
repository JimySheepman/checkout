package models

import (
	"encoding/json"
)

type AddItemServiceRequest struct {
	ItemID     int     `json:"itemId"`
	CategoryID int     `json:"categoryId"`
	SellerID   int     `json:"sellerId"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}

type AddVasItemToItemServiceRequest struct {
	ItemID     int     `json:"itemId"`
	VasItemID  int     `json:"vasItemId"`
	CategoryID int     `json:"categoryId"`
	SellerID   int     `json:"sellerId"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}

type DisplayCartServiceResponse struct {
	Items              []ItemServiceResponse `json:"items"`
	TotalPrice         float64               `json:"totalPrice"`
	AppliedPromotionID int                   `json:"appliedPromotionId"`
	TotalDiscount      float64               `json:"totalDiscount"`
}

func (r *DisplayCartServiceResponse) ToString() string {
	res, _ := json.Marshal(r)
	return string(res)
}

type ItemServiceResponse struct {
	ItemID     int                      `json:"itemId"`
	CategoryID int                      `json:"categoryId"`
	SellerID   int                      `json:"sellerId"`
	Price      float64                  `json:"price"`
	Quantity   int                      `json:"quantity"`
	VasItems   []VasItemServiceResponse `json:"vasItems"`
}

type VasItemServiceResponse struct {
	VasItemID  int     `json:"vasItemId"`
	CategoryID int     `json:"categoryId"`
	SellerID   int     `json:"sellerId"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}

type PromotionServiceResponse struct {
	AppliedPromotionID int     `json:"appliedPromotionId"`
	TotalDiscount      float64 `json:"totalDiscount"`
	TotalPrice         float64 `json:"totalPrice"`
}
