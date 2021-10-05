package controller

import (
	"awesomeProject1/controller"
	mock_services "awesomeProject1/mock/services"
	"awesomeProject1/models"
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetProductById_ProductUnavailable(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	id:=2
	MockServices :=mock_services.NewMockProductService(ctrl)
	MockServices.EXPECT().GetProductById(id).Return(models.Product{},0)
	ProductController:=controller.Initialise(MockServices)
	req, _ := http.NewRequest("GET", "/products/2", nil)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/products/{id}",ProductController.GetProductById ).Methods("GET")
	r.ServeHTTP(rr, req)
	StringResponse:=`{"message":"Product Not available"}`
	str:=strings.TrimSpace(rr.Body.String())
	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 400,rr.Code)
}

func TestGetProductById_ProductAvailable(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	id:=2
	MockServices :=mock_services.NewMockProductService(ctrl)
	product:=models.Product{2,"wickets","wood",256,3}
	MockServices.EXPECT().GetProductById(id).Return(product,1)
	ProductController:=controller.Initialise(MockServices)
	req, _ := http.NewRequest("GET", "/products/2", nil)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/products/{id}",ProductController.GetProductById ).Methods("GET")
	r.ServeHTTP(rr, req)
	StringResponse:=`{"id":2,"name":"wickets","description":"wood","price":256,"quantity":3}`
	str:=strings.TrimSpace(rr.Body.String())
	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}

func TestCreateProduct(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()

	InputProduct:=models.Product{Id:0,Name:"wickets",Description :"wood",Price: 256,Quantity: 3}
	OutputProduct:=models.Product{Id:1,Name:"wickets",Description :"wood",Price: 256,Quantity: 3}
	BodyProduct := []byte(`{"name":"wickets","description":"wood","price":256,"quantity":3}`)

	req, _ := http.NewRequest("PUT", "/products", bytes.NewBuffer(BodyProduct))
	MockServices :=mock_services.NewMockProductService(ctrl)
	MockServices.EXPECT().AddNewProduct(InputProduct).Return(OutputProduct)
	ProductController:=controller.Initialise(MockServices)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/products",ProductController.CreateProduct ).Methods("PUT")
	r.ServeHTTP(rr, req)

	StringResponse:=`{"id":1,"name":"wickets","description":"wood","price":256,"quantity":3}`
	str:=strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}

func TestIncreaseQuantity(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()

	BodyProduct := []byte(`{"quantity":3}`)
    InputQuantity:=models.Product{Quantity:3}
	req, _ := http.NewRequest("PUT", "/products/update/1", bytes.NewBuffer(BodyProduct))
	MockServices :=mock_services.NewMockProductService(ctrl)
	MockServices.EXPECT().IncreaseProductQuantity(InputQuantity,1).Return()

	ProductController:=controller.Initialise(MockServices)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/products/update/{id}",ProductController.IncreaseQuantity ).Methods("PUT")
	r.ServeHTTP(rr, req)

	StringResponse:=`{"message":"Product Update Successful"}`
	str:=strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}


func TestGetAllProduct(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()

	MockServices :=mock_services.NewMockProductService(ctrl)
	ReturnedProduct:=make([]models.Product,0)
	ReturnedProduct=append(ReturnedProduct,models.Product{Id: 1,Name: "bat",Description: "wood",Price: 2000,Quantity: 10})
	MockServices.EXPECT().GetAllProduct().Return(ReturnedProduct)

	ProductController:=controller.Initialise(MockServices)

	req, _ := http.NewRequest("GET", "/products", nil)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/products",ProductController.GetAllProduct ).Methods("GET")
	r.ServeHTTP(rr, req)

	StringResponse:=`[{"id":1,"name":"bat","description":"wood","price":2000,"quantity":10}]`
	str:=strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}


func TestGetAvailableProduct(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()

	MockServices :=mock_services.NewMockProductService(ctrl)
	ReturnedProduct:=make([]models.Product,0)
	ReturnedProduct=append(ReturnedProduct,models.Product{Id: 1,Name: "bat",Description: "wood",Price: 2000,Quantity: 10})
	MockServices.EXPECT().GetAvailableProduct().Return(ReturnedProduct)

	ProductController:=controller.Initialise(MockServices)

	req, _ := http.NewRequest("GET", "/products", nil)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/products",ProductController.GetAvailableProduct).Methods("GET")
	r.ServeHTTP(rr, req)

	StringResponse:=`[{"id":1,"name":"bat","description":"wood","price":2000,"quantity":10}]`
	str:=strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}


func TestBuyProduct_ProductUnavailable(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()

	MockServices :=mock_services.NewMockProductService(ctrl)
	id:=1
	value:=models.Product{Quantity: 20}
	Quantity := []byte(`{"quantity":20}`)


	MockServices.EXPECT().ReduceProductQuantity(value,id).Return(1)

	ProductController:=controller.Initialise(MockServices)

	req, _ := http.NewRequest("PUT", "/products/purchase/1", bytes.NewBuffer(Quantity))
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/products/purchase/{id}",ProductController.BuyProduct).Methods("PUT")
	r.ServeHTTP(rr, req)

	StringResponse:=`{"message":"Product Not available"}`
	str:=strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 400,rr.Code)
}


func TestBuyProduct_ProductInsufficient(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()

	MockServices :=mock_services.NewMockProductService(ctrl)
	id:=1
	value:=models.Product{Quantity: 20}
	Quantity := []byte(`{"quantity":20}`)


	MockServices.EXPECT().ReduceProductQuantity(value,id).Return(2)

	ProductController:=controller.Initialise(MockServices)

	req, _ := http.NewRequest("PUT", "/products/purchase/1", bytes.NewBuffer(Quantity))
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/products/purchase/{id}",ProductController.BuyProduct).Methods("PUT")
	r.ServeHTTP(rr, req)

	StringResponse:=`{"message":"sufficient quantity not available"}`
	str:=strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 400,rr.Code)
}

func TestBuyProduct_ProductAvailable(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()

	MockServices :=mock_services.NewMockProductService(ctrl)
	id:=1
	value:=models.Product{Quantity: 20}
	Quantity := []byte(`{"quantity":20}`)


	MockServices.EXPECT().ReduceProductQuantity(value,id).Return(3)

	ProductController:=controller.Initialise(MockServices)

	req, _ := http.NewRequest("PUT", "/products/purchase/1", bytes.NewBuffer(Quantity))
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/products/purchase/{id}",ProductController.BuyProduct).Methods("PUT")
	r.ServeHTTP(rr, req)

	StringResponse:=`{"message":"Product purchase successful"}`
	str:=strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}



