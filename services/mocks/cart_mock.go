// Code generated by MockGen. DO NOT EDIT.
// Source: cart.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	domain "checkout-case/domain"
	models "checkout-case/models"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockcartRepository is a mock of cartRepository interface.
type MockcartRepository struct {
	ctrl     *gomock.Controller
	recorder *MockcartRepositoryMockRecorder
}

// MockcartRepositoryMockRecorder is the mock recorder for MockcartRepository.
type MockcartRepositoryMockRecorder struct {
	mock *MockcartRepository
}

// NewMockcartRepository creates a new mock instance.
func NewMockcartRepository(ctrl *gomock.Controller) *MockcartRepository {
	mock := &MockcartRepository{ctrl: ctrl}
	mock.recorder = &MockcartRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcartRepository) EXPECT() *MockcartRepositoryMockRecorder {
	return m.recorder
}

// AddItem mocks base method.
func (m *MockcartRepository) AddItem(item *domain.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddItem", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddItem indicates an expected call of AddItem.
func (mr *MockcartRepositoryMockRecorder) AddItem(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddItem", reflect.TypeOf((*MockcartRepository)(nil).AddItem), item)
}

// AddVasItemToItemByItemID mocks base method.
func (m *MockcartRepository) AddVasItemToItemByItemID(itemId primitive.ObjectID, vasItem *domain.VasItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddVasItemToItemByItemID", itemId, vasItem)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddVasItemToItemByItemID indicates an expected call of AddVasItemToItemByItemID.
func (mr *MockcartRepositoryMockRecorder) AddVasItemToItemByItemID(itemId, vasItem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVasItemToItemByItemID", reflect.TypeOf((*MockcartRepository)(nil).AddVasItemToItemByItemID), itemId, vasItem)
}

// FindItemByItemIdFromCart mocks base method.
func (m *MockcartRepository) FindItemByItemIdFromCart(itemId int) (*domain.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindItemByItemIdFromCart", itemId)
	ret0, _ := ret[0].(*domain.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindItemByItemIdFromCart indicates an expected call of FindItemByItemIdFromCart.
func (mr *MockcartRepositoryMockRecorder) FindItemByItemIdFromCart(itemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindItemByItemIdFromCart", reflect.TypeOf((*MockcartRepository)(nil).FindItemByItemIdFromCart), itemId)
}

// FindVasItemByVasItemIdFromItem mocks base method.
func (m *MockcartRepository) FindVasItemByVasItemIdFromItem(vasItemId int) (*domain.VasItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindVasItemByVasItemIdFromItem", vasItemId)
	ret0, _ := ret[0].(*domain.VasItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindVasItemByVasItemIdFromItem indicates an expected call of FindVasItemByVasItemIdFromItem.
func (mr *MockcartRepositoryMockRecorder) FindVasItemByVasItemIdFromItem(vasItemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindVasItemByVasItemIdFromItem", reflect.TypeOf((*MockcartRepository)(nil).FindVasItemByVasItemIdFromItem), vasItemId)
}

// GetCart mocks base method.
func (m *MockcartRepository) GetCart() (*domain.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCart")
	ret0, _ := ret[0].(*domain.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCart indicates an expected call of GetCart.
func (mr *MockcartRepositoryMockRecorder) GetCart() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCart", reflect.TypeOf((*MockcartRepository)(nil).GetCart))
}

// RemoveItem mocks base method.
func (m *MockcartRepository) RemoveItem(item *domain.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveItem", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveItem indicates an expected call of RemoveItem.
func (mr *MockcartRepositoryMockRecorder) RemoveItem(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveItem", reflect.TypeOf((*MockcartRepository)(nil).RemoveItem), item)
}

// ResetCart mocks base method.
func (m *MockcartRepository) ResetCart() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetCart")
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetCart indicates an expected call of ResetCart.
func (mr *MockcartRepositoryMockRecorder) ResetCart() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetCart", reflect.TypeOf((*MockcartRepository)(nil).ResetCart))
}

// UpdateItemQuantity mocks base method.
func (m *MockcartRepository) UpdateItemQuantity(item *domain.Item, req *models.AddItemServiceRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItemQuantity", item, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateItemQuantity indicates an expected call of UpdateItemQuantity.
func (mr *MockcartRepositoryMockRecorder) UpdateItemQuantity(item, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItemQuantity", reflect.TypeOf((*MockcartRepository)(nil).UpdateItemQuantity), item, req)
}

// UpdateVasItemQuantity mocks base method.
func (m *MockcartRepository) UpdateVasItemQuantity(item *domain.Item, vasItem *domain.VasItem, req *models.AddVasItemToItemServiceRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVasItemQuantity", item, vasItem, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVasItemQuantity indicates an expected call of UpdateVasItemQuantity.
func (mr *MockcartRepositoryMockRecorder) UpdateVasItemQuantity(item, vasItem, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVasItemQuantity", reflect.TypeOf((*MockcartRepository)(nil).UpdateVasItemQuantity), item, vasItem, req)
}

// MockpromotionServiceClient is a mock of promotionServiceClient interface.
type MockpromotionServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockpromotionServiceClientMockRecorder
}

// MockpromotionServiceClientMockRecorder is the mock recorder for MockpromotionServiceClient.
type MockpromotionServiceClientMockRecorder struct {
	mock *MockpromotionServiceClient
}

// NewMockpromotionServiceClient creates a new mock instance.
func NewMockpromotionServiceClient(ctrl *gomock.Controller) *MockpromotionServiceClient {
	mock := &MockpromotionServiceClient{ctrl: ctrl}
	mock.recorder = &MockpromotionServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpromotionServiceClient) EXPECT() *MockpromotionServiceClientMockRecorder {
	return m.recorder
}

// FindBestPromotion mocks base method.
func (m *MockpromotionServiceClient) FindBestPromotion(ctx context.Context, cart *domain.Cart) (*models.PromotionServiceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBestPromotion", ctx, cart)
	ret0, _ := ret[0].(*models.PromotionServiceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBestPromotion indicates an expected call of FindBestPromotion.
func (mr *MockpromotionServiceClientMockRecorder) FindBestPromotion(ctx, cart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBestPromotion", reflect.TypeOf((*MockpromotionServiceClient)(nil).FindBestPromotion), ctx, cart)
}
