package models

import "encoding/json"

type CommandRequest struct {
	Command string `json:"command"`
}

type AddItemCommandRequest struct {
	ItemID     int     `json:"itemId"`
	CategoryID int     `json:"categoryId"`
	SellerID   int     `json:"sellerId"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}

type AddVasItemToItemCommandRequest struct {
	ItemID     int     `json:"itemId"`
	VasItemID  int     `json:"vasItemId"`
	CategoryID int     `json:"categoryId"`
	SellerID   int     `json:"sellerId"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}

type RemoveItemCommandRequest struct {
	ItemID int `json:"itemId"`
}

type GenericCommandResponse struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

type DisplayCartCommandResponse struct {
	Result  bool                              `json:"result"`
	Message DisplayCartMessageCommandResponse `json:"message"`
}

type DisplayCartMessageCommandResponse struct {
	Items              []ItemCommandResponse `json:"items"`
	TotalPrice         float64               `json:"totalPrice"`
	AppliedPromotionID int                   `json:"appliedPromotionId"`
	TotalDiscount      float64               `json:"totalDiscount"`
}

type ItemCommandResponse struct {
	ItemID     int                      `json:"itemId"`
	CategoryID int                      `json:"categoryId"`
	SellerID   int                      `json:"sellerId"`
	Price      float64                  `json:"price"`
	Quantity   int                      `json:"quantity"`
	VasItems   []VasItemCommandResponse `json:"vasItems"`
}

type VasItemCommandResponse struct {
	VasItemID  int     `json:"vasItemId"`
	CategoryID int     `json:"categoryId"`
	SellerID   int     `json:"sellerId"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}

type AddItemRestRequest struct {
	ItemID     int     `json:"itemId"`
	CategoryID int     `json:"categoryId"`
	SellerID   int     `json:"sellerId"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}

type AddVasItemToItemRestRequest struct {
	ItemID     int     `json:"itemId"`
	VasItemID  int     `json:"vasItemId"`
	CategoryID int     `json:"categoryId"`
	SellerID   int     `json:"sellerId"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}

type DisplayCartRestResponse struct {
	Result  bool                           `json:"result"`
	Message DisplayCartMessageRestResponse `json:"message"`
}

type DisplayCartMessageRestResponse struct {
	Items              []ItemRestResponse `json:"items"`
	TotalPrice         float64            `json:"totalPrice"`
	AppliedPromotionID int                `json:"appliedPromotionId"`
	TotalDiscount      float64            `json:"totalDiscount"`
}

type ItemRestResponse struct {
	ItemID     int                   `json:"itemId"`
	CategoryID int                   `json:"categoryId"`
	SellerID   int                   `json:"sellerId"`
	Price      float64               `json:"price"`
	Quantity   int                   `json:"quantity"`
	VasItems   []VasItemRestResponse `json:"vasItems"`
}

type VasItemRestResponse struct {
	VasItemID  int     `json:"vasItemId"`
	CategoryID int     `json:"categoryId"`
	SellerID   int     `json:"sellerId"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}

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

func (r DisplayCartServiceResponse) ToString() string {
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
