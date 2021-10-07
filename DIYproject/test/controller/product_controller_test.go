package controller

import (
	"awesomeProject1/controller"
	"awesomeProject1/errors"
	"awesomeProject1/models"
	"awesomeProject1/test/mock/services"
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGetProductById_ProductUnavailable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := 2
	MockServices := mock_services.NewMockProductService(ctrl)
	MockServices.EXPECT().GetProductById(id).Return(models.Product{}, errors.ProductNotFound)
	ProductController := controller.Initialise(MockServices)
	req, _ := http.NewRequest("GET", "/products/2", nil)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/products/{id}", ProductController.GetProductById).Methods("GET")
	r.ServeHTTP(rr, req)
	StringResponse := `{"message":"Product Not available"}`
	str := strings.TrimSpace(rr.Body.String())
	assert.Equal(t, StringResponse, str)
	assert.Equal(t, 200, rr.Code)
}

func TestGetProductById_ProductAvailable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := 2
	MockServices := mock_services.NewMockProductService(ctrl)
	product := models.Product{2, "wickets", "wood", 256, 3}
	MockServices.EXPECT().GetProductById(id).Return(product, nil)
	ProductController := controller.Initialise(MockServices)
	req, _ := http.NewRequest("GET", "/products/2", nil)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/products/{id}", ProductController.GetProductById).Methods("GET")
	r.ServeHTTP(rr, req)
	StringResponse := `{"id":2,"name":"wickets","description":"wood","price":256,"quantity":3}`
	str := strings.TrimSpace(rr.Body.String())
	assert.Equal(t, StringResponse, str)
	assert.Equal(t, 200, rr.Code)
}

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	InputProduct := models.Product{Id: 0, Name: "wickets", Description: "wood", Price: 256, Quantity: 3}
	OutputProduct := models.Product{Id: 1, Name: "wickets", Description: "wood", Price: 256, Quantity: 3}
	BodyProduct := []byte(`{"name":"wickets","description":"wood","price":256,"quantity":3}`)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(BodyProduct))
	MockServices := mock_services.NewMockProductService(ctrl)
	MockServices.EXPECT().AddNewProduct(InputProduct).Return(OutputProduct)
	ProductController := controller.Initialise(MockServices)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/products", ProductController.CreateProduct).Methods("POST")
	r.ServeHTTP(rr, req)

	StringResponse := `{"id":1,"name":"wickets","description":"wood","price":256,"quantity":3}`
	str := strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse, str)
	assert.Equal(t, 200, rr.Code)
}

func TestIncreaseQuantity_ProductAvailable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	BodyProduct := []byte(`{"quantity":3}`)
	InputQuantity := models.Product{Quantity: 3}
	req, _ := http.NewRequest("PUT", "/products/update/1", bytes.NewBuffer(BodyProduct))
	MockServices := mock_services.NewMockProductService(ctrl)
	MockServices.EXPECT().IncreaseProductQuantity(InputQuantity.Quantity, 1).Return(nil)

	ProductController := controller.Initialise(MockServices)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/products/update/{id}", ProductController.IncreaseQuantity).Methods("PUT")
	r.ServeHTTP(rr, req)

	StringResponse := `{"message":"product update successful"}`
	str := strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse, str)
	assert.Equal(t, 200, rr.Code)
}

func TestIncreaseQuantity_ProductUnAvailable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	BodyProduct := []byte(`{"quantity":3}`)
	InputQuantity := models.Product{Quantity: 3}
	req, _ := http.NewRequest("PUT", "/products/update/1", bytes.NewBuffer(BodyProduct))
	MockServices := mock_services.NewMockProductService(ctrl)
	MockServices.EXPECT().IncreaseProductQuantity(InputQuantity.Quantity, 1).Return(errors.ProductNotFound)

	ProductController := controller.Initialise(MockServices)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/products/update/{id}", ProductController.IncreaseQuantity).Methods("PUT")
	r.ServeHTTP(rr, req)

	StringResponse := `{"message":"product does not exist"}`
	str := strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse, str)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetAllProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockServices := mock_services.NewMockProductService(ctrl)
	ReturnedProduct := make([]models.Product, 0)
	ReturnedProduct = append(ReturnedProduct, models.Product{Id: 1, Name: "bat", Description: "wood", Price: 2000, Quantity: 10})
	MockServices.EXPECT().GetAllProducts().Return(ReturnedProduct)

	ProductController := controller.Initialise(MockServices)

	req, _ := http.NewRequest("GET", "/products", nil)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/products", ProductController.GetAllProducts).Methods("GET")
	r.ServeHTTP(rr, req)

	StringResponse := `[{"id":1,"name":"bat","description":"wood","price":2000,"quantity":10}]`
	str := strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse, str)
	assert.Equal(t, 200, rr.Code)
}

func TestGetAvailableProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockServices := mock_services.NewMockProductService(ctrl)
	ReturnedProduct := make([]models.Product, 0)
	ReturnedProduct = append(ReturnedProduct, models.Product{Id: 1, Name: "bat", Description: "wood", Price: 2000, Quantity: 10})
	MockServices.EXPECT().GetAvailableProducts().Return(ReturnedProduct)

	ProductController := controller.Initialise(MockServices)

	req, _ := http.NewRequest("GET", "/products", nil)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/products", ProductController.GetAvailableProducts).Methods("GET")
	r.ServeHTTP(rr, req)

	StringResponse := `[{"id":1,"name":"bat","description":"wood","price":2000,"quantity":10}]`
	str := strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse, str)
	assert.Equal(t, 200, rr.Code)
}

func TestBuyProduct_ProductUnavailable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockServices := mock_services.NewMockProductService(ctrl)
	id := 1
	value := models.Product{Quantity: 20}
	Quantity := []byte(`{"quantity":20}`)

	MockServices.EXPECT().ReduceProductQuantity(value.Quantity, id).Return(errors.ProductNotFound)

	ProductController := controller.Initialise(MockServices)

	req, _ := http.NewRequest("PUT", "/products/purchase/1", bytes.NewBuffer(Quantity))
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/products/purchase/{id}", ProductController.BuyProduct).Methods("PUT")
	r.ServeHTTP(rr, req)

	StringResponse := `{"message":"product not available"}`
	str := strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse, str)
	assert.Equal(t, 200, rr.Code)
}

func TestBuyProduct_ProductInsufficient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockServices := mock_services.NewMockProductService(ctrl)
	id := 1
	value := models.Product{Quantity: 20}
	Quantity := []byte(`{"quantity":20}`)

	MockServices.EXPECT().ReduceProductQuantity(value.Quantity, id).Return(errors.InsufficientProduct)

	ProductController := controller.Initialise(MockServices)

	req, _ := http.NewRequest("PUT", "/products/purchase/1", bytes.NewBuffer(Quantity))
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/products/purchase/{id}", ProductController.BuyProduct).Methods("PUT")
	r.ServeHTTP(rr, req)

	StringResponse := `{"message":"sufficient quantity not available"}`
	str := strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse, str)
	assert.Equal(t, 200, rr.Code)
}

func TestBuyProduct_ProductAvailable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockServices := mock_services.NewMockProductService(ctrl)
	id := 1
	value := models.Product{Quantity: 20}
	Quantity := []byte(`{"quantity":20}`)

	MockServices.EXPECT().ReduceProductQuantity(value.Quantity, id).Return(nil)
	CurrentTime := time.Now()
	MockServices.EXPECT().GetCurrentTime().Return(CurrentTime)
	MockServices.EXPECT().AddNewSalesRecord(models.SalesRecord{Id: 0, ProductId: id, QuantitySold: value.Quantity, SalesTime: CurrentTime}).Return()
	ProductController := controller.Initialise(MockServices)

	req, _ := http.NewRequest("PUT", "/products/purchase/1", bytes.NewBuffer(Quantity))
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/products/purchase/{id}", ProductController.BuyProduct).Methods("PUT")
	r.ServeHTTP(rr, req)

	StringResponse := `{"message":"product purchase successful"}`
	str := strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse, str)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGeTop5Product_NoPurchaseDone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockServices := mock_services.NewMockProductService(ctrl)

	Top5ProductZeroLength := make([]models.TopProductResponse, 0)
	MockServices.EXPECT().GetTop5Product().Return(Top5ProductZeroLength)
	ProductController := controller.Initialise(MockServices)

	req, _ := http.NewRequest("GET", "/products/topsold", nil)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/products/topsold", ProductController.GetTop5Product).Methods("GET")
	r.ServeHTTP(rr, req)

	StringResponse := `{"message":"no purchase in last one hour"}`
	str := strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse, str)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGeTop5Product_PurchaseDone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockServices := mock_services.NewMockProductService(ctrl)

	Top5ProductNonZeroLength := make([]models.TopProductResponse, 0)
	Top5ProductNonZeroLength = append(Top5ProductNonZeroLength, models.TopProductResponse{ProductId: 1, QuantitySold: 22})
	MockServices.EXPECT().GetTop5Product().Return(Top5ProductNonZeroLength)
	ProductController := controller.Initialise(MockServices)

	req, _ := http.NewRequest("GET", "/products/topsold", nil)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/products/topsold", ProductController.GetTop5Product).Methods("GET")
	r.ServeHTTP(rr, req)

	StringResponse := `[{"productid":1,"quantitysold":22}]`
	str := strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse, str)
	assert.Equal(t, http.StatusOK, rr.Code)
}
