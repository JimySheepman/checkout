package models

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
