package models

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
