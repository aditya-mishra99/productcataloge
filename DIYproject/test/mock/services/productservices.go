// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/productservices.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	models "awesomeProject1/models"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockProductService is a mock of ProductService interface.
type MockProductService struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceMockRecorder
}

// MockProductServiceMockRecorder is the mock recorder for MockProductService.
type MockProductServiceMockRecorder struct {
	mock *MockProductService
}

// NewMockProductService creates a new mock instance.
func NewMockProductService(ctrl *gomock.Controller) *MockProductService {
	mock := &MockProductService{ctrl: ctrl}
	mock.recorder = &MockProductServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductService) EXPECT() *MockProductServiceMockRecorder {
	return m.recorder
}

// AddNewProduct mocks base method.
func (m *MockProductService) AddNewProduct(product models.Product) models.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewProduct", product)
	ret0, _ := ret[0].(models.Product)
	return ret0
}

// AddNewProduct indicates an expected call of AddNewProduct.
func (mr *MockProductServiceMockRecorder) AddNewProduct(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewProduct", reflect.TypeOf((*MockProductService)(nil).AddNewProduct), product)
}

// AddNewSalesRecord mocks base method.
func (m *MockProductService) AddNewSalesRecord(SaleRecord models.SalesRecord) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddNewSalesRecord", SaleRecord)
}

// AddNewSalesRecord indicates an expected call of AddNewSalesRecord.
func (mr *MockProductServiceMockRecorder) AddNewSalesRecord(SaleRecord interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewSalesRecord", reflect.TypeOf((*MockProductService)(nil).AddNewSalesRecord), SaleRecord)
}

// GetAllProducts mocks base method.
func (m *MockProductService) GetAllProducts() []models.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProducts")
	ret0, _ := ret[0].([]models.Product)
	return ret0
}

// GetAllProducts indicates an expected call of GetAllProducts.
func (mr *MockProductServiceMockRecorder) GetAllProducts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProducts", reflect.TypeOf((*MockProductService)(nil).GetAllProducts))
}

// GetAvailableProducts mocks base method.
func (m *MockProductService) GetAvailableProducts() []models.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableProducts")
	ret0, _ := ret[0].([]models.Product)
	return ret0
}

// GetAvailableProducts indicates an expected call of GetAvailableProducts.
func (mr *MockProductServiceMockRecorder) GetAvailableProducts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableProducts", reflect.TypeOf((*MockProductService)(nil).GetAvailableProducts))
}

// GetCurrentTime mocks base method.
func (m *MockProductService) GetCurrentTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetCurrentTime indicates an expected call of GetCurrentTime.
func (mr *MockProductServiceMockRecorder) GetCurrentTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentTime", reflect.TypeOf((*MockProductService)(nil).GetCurrentTime))
}

// GetProductById mocks base method.
func (m *MockProductService) GetProductById(id int) (models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductById", id)
	ret0, _ := ret[0].(models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductById indicates an expected call of GetProductById.
func (mr *MockProductServiceMockRecorder) GetProductById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductById", reflect.TypeOf((*MockProductService)(nil).GetProductById), id)
}

// GetTop5Product mocks base method.
func (m *MockProductService) GetTop5Product() []models.TopProductResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTop5Product")
	ret0, _ := ret[0].([]models.TopProductResponse)
	return ret0
}

// GetTop5Product indicates an expected call of GetTop5Product.
func (mr *MockProductServiceMockRecorder) GetTop5Product() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTop5Product", reflect.TypeOf((*MockProductService)(nil).GetTop5Product))
}

// IncreaseProductQuantity mocks base method.
func (m *MockProductService) IncreaseProductQuantity(quantity, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncreaseProductQuantity", quantity, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncreaseProductQuantity indicates an expected call of IncreaseProductQuantity.
func (mr *MockProductServiceMockRecorder) IncreaseProductQuantity(quantity, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseProductQuantity", reflect.TypeOf((*MockProductService)(nil).IncreaseProductQuantity), quantity, id)
}

// ReduceProductQuantity mocks base method.
func (m *MockProductService) ReduceProductQuantity(quantity, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReduceProductQuantity", quantity, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReduceProductQuantity indicates an expected call of ReduceProductQuantity.
func (mr *MockProductServiceMockRecorder) ReduceProductQuantity(quantity, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReduceProductQuantity", reflect.TypeOf((*MockProductService)(nil).ReduceProductQuantity), quantity, id)
}
