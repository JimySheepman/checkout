package cart

import (
	"checkout-case/internal/core/domain"
	"checkout-case/internal/core/models"
	"checkout-case/internal/core/port"
	rule_item "checkout-case/internal/core/service/cart/chain/item"
	rule_vasitem "checkout-case/internal/core/service/cart/chain/vasitem"
	"checkout-case/pkg/logger"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	vasItemCategoryID     = 3242
	vasItemSellerID       = 5003
	DigitalItemCategoryID = 7889
)

type cartService struct {
	cartRepository         port.CartRepository
	promotionServiceClient port.PromotionServiceClient
}

func NewCartService(cartRepository port.CartRepository, promotionServiceClient port.PromotionServiceClient) *cartService {
	return &cartService{
		cartRepository:         cartRepository,
		promotionServiceClient: promotionServiceClient,
	}
}

func (s *cartService) AddItemToCart(ctx context.Context, req *models.AddItemServiceRequest) error {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("AddItemToCart service is start")

	if err := s.addItemToCartValidation(req); err != nil {
		return fmt.Errorf("add item to cart validation: %w", err)
	}
	l.Info("addItemToCartValidation method in AddItemToCart completed")

	cart, err := s.cartRepository.GetCart()
	if err != nil {
		return fmt.Errorf("get cart: %w", err)
	}
	l.Info("GetCart method in AddItemToCart completed")

	// There are two possibilities here, I need to update the quantity if the
	// item is there, the other is to add the item directly. For this, I keep
	// a flag and update and return if this object is present.
	tmpCart, tmpItem, ok := populateTempCartForItem(cart, req)

	if !s.addItemToCartRuleChain(ctx, tmpCart) {
		return fmt.Errorf("could not add item to cart")
	}
	l.Info("addItemToCartRuleChain method in AddItemToCart completed")

	if ok {
		if err := s.sameItemProcess(ctx, tmpItem, req); err != nil {
			return fmt.Errorf("same item process: %w", err)
		}

		return nil
	}

	if err := s.cartRepository.AddItem(populateAddItemServiceRequestToItem(req)); err != nil {
		return fmt.Errorf("could not add item to cart error: %w", err)
	}
	l.Info("AddItem method in AddItemToCart completed")
	l.Info("AddItemToCart done")

	return nil
}

func (s *cartService) addItemToCartValidation(req *models.AddItemServiceRequest) error {
	if req.CategoryID == vasItemCategoryID {
		return fmt.Errorf("wrong categoryID")
	}

	if req.SellerID == vasItemSellerID {
		return fmt.Errorf("wrong sellerID")
	}

	return nil
}

func (s *cartService) sameItemProcess(ctx context.Context, item *domain.Item, req *models.AddItemServiceRequest) error {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("sameItemProcess service is start")

	if err := s.cartRepository.UpdateItemQuantity(item, req); err != nil {
		return fmt.Errorf("could not add item to cart error: %w", err)
	}
	l.Info("UpdateItemQuantity method in sameItemProcess completed")
	l.Info("sameItemProcess done")

	return nil
}

func (s *cartService) addItemToCartRuleChain(ctx context.Context, cart *domain.Cart) bool {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("Item Rule Chain service is start")

	for _, rule := range rule_item.GetRulers() {
		if err := rule.IsAddItemValid(ctx, cart); err != nil {
			l.Errorf("addItemToCartRuleChain: %v", err)
			return false
		}
	}

	return true
}

func populateAddItemServiceRequestToItem(req *models.AddItemServiceRequest) *domain.Item {
	now := time.Now()

	item := &domain.Item{
		ID:         primitive.NewObjectID().String(),
		ItemId:     req.ItemID,
		CategoryId: req.CategoryID,
		SellerId:   req.SellerID,
		CreatedAt:  now,
		UpdatedAt:  now,
		Price:      req.Price,
		Quantity:   req.Quantity,
	}

	if req.CategoryID == DigitalItemCategoryID {
		item.Type = domain.DigitalItem
	} else {
		item.Type = domain.DefaultItem
		item.VasItems = []*domain.VasItem{}
	}

	return item
}

func populateTempCartForItem(cart *domain.Cart, req *models.AddItemServiceRequest) (*domain.Cart, *domain.Item, bool) {
	var (
		flag     bool
		tmpItems []*domain.Item
		tmpItem  *domain.Item
	)

	for _, item := range cart.Items {
		if item.ItemId == req.ItemID {
			flag = true
			item.Quantity += req.Quantity
			cart.TotalPrice += float64(req.Quantity) * req.Price
			tmpItem = item
		}

		tmpItems = append(tmpItems, item)
	}

	if !flag {
		i := &domain.Item{
			ItemId:     req.ItemID,
			CategoryId: req.CategoryID,
			SellerId:   req.SellerID,
			Price:      req.Price,
			Quantity:   req.Quantity,
		}

		cart.TotalPrice += float64(i.Quantity) * i.Price

		tmpItems = append(tmpItems, i)
	}

	c := &domain.Cart{
		Items:      tmpItems,
		TotalPrice: cart.TotalPrice,
	}

	return c, tmpItem, flag
}

func (s *cartService) AddVasItemToItem(ctx context.Context, req *models.AddVasItemToItemServiceRequest) error {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("AddVasItemToItem service is start")

	item, err := s.cartRepository.FindItemByItemIdFromCart(req.ItemID)
	if err != nil {
		return fmt.Errorf("find item by itemId from cart: %w", err)
	}
	l.Info("FindItemByItemIdFromCart method in AddVasItemToItem completed")

	if err := s.addVasItemToItemValidation(item, req); err != nil {
		return fmt.Errorf("add vasItem to item validation: %w", err)
	}
	l.Info("addVasItemToItemValidation method in AddVasItemToItem completed")

	cart, err := s.cartRepository.GetCart()
	if err != nil {
		return fmt.Errorf("same vasItem process: %w", err)
	}
	l.Info("GetCart method in AddVasItemToItem completed")

	// There are two possibilities here, I need to update the quantity if the
	// vasItem is there, the other is to add the vasItem directly. For this,
	// I keep a flag and update and return if this object is present.
	tmpItem, tmpVasItem, totalPrice, ok := populateTempCartForVasItem(cart, item, req)

	if !s.addVasItemToItemRuleChain(ctx, tmpItem, totalPrice) {
		return fmt.Errorf("could not add item to cart")
	}
	l.Info("addVasItemToItemRuleChain method in AddVasItemToItem completed")

	if ok {
		l.Info("sameVasItemProcess service is start")

		if err := s.cartRepository.UpdateVasItemQuantity(item, tmpVasItem, req); err != nil {
			return fmt.Errorf("could not update item to cart error: %w", err)
		}
		l.Info("UpdateVasItemQuantity method in sameVasItemProcess completed")
		l.Info("sameVasItemProcess done")

		return nil
	}

	if err := s.cartRepository.AddVasItemToItemByItemID(item.ID, populateAddVasItemToItemServiceRequestToVasItem(req)); err != nil {
		return fmt.Errorf("could not add item to cart error: %w", err)
	}
	l.Info("AddVasItemToItemByItemID method in AddVasItemToItem completed")
	l.Info("AddVasItemToItem done")

	return nil
}

func (s *cartService) addVasItemToItemValidation(item *domain.Item, req *models.AddVasItemToItemServiceRequest) error {
	if req.Price >= item.Price {
		return fmt.Errorf("wrong item price for vasItem")
	}

	if item.Type == domain.DigitalItem {
		return fmt.Errorf("wrong item type for vasItem")
	}

	if req.CategoryID != vasItemCategoryID {
		return fmt.Errorf("vasItemID is not correct")
	}

	return nil
}

func (s *cartService) addVasItemToItemRuleChain(ctx context.Context, item *domain.Item, totalPrice float64) bool {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("VasItem Rule Chain service is start")

	for _, rule := range rule_vasitem.GetVasRulers() {
		if err := rule.IsAddVasItemValid(ctx, item, totalPrice); err != nil {
			l.Errorf("addVasItemToItemRuleChain: %v", err)
			return false
		}
	}

	return true
}

func populateTempCartForVasItem(cart *domain.Cart, item *domain.Item, req *models.AddVasItemToItemServiceRequest) (*domain.Item, *domain.VasItem, float64, bool) {
	var (
		flag        bool
		totalPrice  = cart.TotalPrice
		tmpVasItems []*domain.VasItem
		tmpVasItem  *domain.VasItem
	)

	for _, vasItem := range item.VasItems {
		if vasItem.VasItemId == req.VasItemID {
			flag = true
			vasItem.Quantity += req.Quantity
			totalPrice += float64(req.Quantity) * req.Price
			tmpVasItem = vasItem
		}

		tmpVasItems = append(tmpVasItems, vasItem)
	}

	if !flag {
		vi := &domain.VasItem{
			VasItemId:  req.VasItemID,
			CategoryId: req.CategoryID,
			SellerId:   req.SellerID,
			Price:      req.Price,
			Quantity:   req.Quantity,
		}

		totalPrice += float64(vi.Quantity) * vi.Price
		tmpVasItems = append(tmpVasItems, vi)
	}

	i := &domain.Item{
		ItemId:     item.ItemId,
		CategoryId: item.CategoryId,
		SellerId:   item.SellerId,
		Type:       item.Type,
		Price:      item.Price,
		Quantity:   item.Quantity,
		VasItems:   tmpVasItems,
	}

	return i, tmpVasItem, totalPrice, flag
}

func populateAddVasItemToItemServiceRequestToVasItem(req *models.AddVasItemToItemServiceRequest) *domain.VasItem {
	return &domain.VasItem{
		ID:         primitive.NewObjectID().String(),
		VasItemId:  req.VasItemID,
		CategoryId: req.CategoryID,
		SellerId:   req.SellerID,
		CreatedAt:  time.Now(),
		Price:      req.Price,
		Quantity:   req.Quantity,
	}
}

func (s *cartService) RemoveItemFromCart(ctx context.Context, itemId int) error {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("RemoveItemFromCart service is start")

	item, err := s.cartRepository.FindItemByItemIdFromCart(itemId)
	if err != nil {
		return fmt.Errorf("remove item from cart: %w", err)
	}
	l.Info("FindItemByItemIdFromCart method in RemoveItemFromCart completed")

	if err := s.cartRepository.RemoveItem(item); err != nil {
		return fmt.Errorf("remove item from cart: %w", err)
	}
	l.Info("RemoveItem method in RemoveItemFromCart completed")
	l.Info("RemoveItemFromCart done")

	return nil
}

func (s *cartService) ResetCart(ctx context.Context) error {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("ResetCart service is start")

	if err := s.cartRepository.ResetCart(); err != nil {
		return fmt.Errorf("reset cart: %w", err)
	}
	l.Info("ResetCart method in ResetCart completed")
	l.Info("ResetCart done")

	return nil
}

func (s *cartService) DisplayCart(ctx context.Context) (*models.DisplayCartServiceResponse, error) {
	l := logger.FromCtx(ctx).Sugar()
	l.Info("DisplayCart service is start")

	cart, err := s.cartRepository.GetCart()
	if err != nil {
		return nil, fmt.Errorf("get cart: %w", err)
	}
	l.Info("GetCart method in DisplayCart completed")

	if cart.TotalPrice == 0 {
		return populateDisplayCartResponse(cart, &models.PromotionServiceResponse{}), nil
	}

	bestPromo, err := s.promotionServiceClient.FindBestPromotion(ctx, cart)
	if err != nil {
		return nil, fmt.Errorf("find best promotion: %w", err)
	}
	l.Info("FindBestPromotion method in DisplayCart done")
	l.Info("DisplayCart done")

	return populateDisplayCartResponse(cart, bestPromo), nil
}

func populateDisplayCartResponse(cart *domain.Cart, promotion *models.PromotionServiceResponse) *models.DisplayCartServiceResponse {
	var (
		itemServiceResponse    []models.ItemServiceResponse
		vasItemServiceResponse []models.VasItemServiceResponse
	)

	for _, item := range cart.Items {
		isr := models.ItemServiceResponse{
			ItemID:     item.ItemId,
			CategoryID: item.CategoryId,
			SellerID:   item.SellerId,
			Price:      item.Price,
			Quantity:   item.Quantity,
		}

		for i := 0; i < len(item.VasItems); i++ {
			visr := models.VasItemServiceResponse{
				VasItemID:  item.VasItems[i].VasItemId,
				CategoryID: item.VasItems[i].CategoryId,
				SellerID:   item.VasItems[i].SellerId,
				Price:      item.VasItems[i].Price,
				Quantity:   item.VasItems[i].Quantity,
			}

			vasItemServiceResponse = append(vasItemServiceResponse, visr)
		}

		isr.VasItems = vasItemServiceResponse

		itemServiceResponse = append(itemServiceResponse, isr)

	}

	return &models.DisplayCartServiceResponse{
		Items:              itemServiceResponse,
		TotalPrice:         promotion.TotalPrice,
		AppliedPromotionID: promotion.AppliedPromotionID,
		TotalDiscount:      promotion.TotalDiscount,
	}
}
