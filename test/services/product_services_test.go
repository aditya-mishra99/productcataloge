package services

import (
	mock_inmemoryrepos "awesomeProject1/mock/repos/inmemoryrepos"
	"awesomeProject1/models"
	"awesomeProject1/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProductById_AvailableProduct(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	id:=1
	MockRepo :=mock_inmemoryrepos.NewMockProductRepo(ctrl)
	MockRepo.EXPECT().GetProductById(id).Return(models.Product{
		Id : 1,
		Price: 22,
		Name : "Balls",
		Quantity: 10,
		Description: "adidas",
	})
	MockRepo.EXPECT().CheckProductAvailableById(id).Return(true)
	ProductService:=services.NewProductService(MockRepo)
	product ,error :=ProductService.GetProductById(id)
	assert.NotNil(t, product)
	assert.Equal(t, 1,error)
}

func TestGetProductById_UnAvailableProduct(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	id:=1
	MockRepo :=mock_inmemoryrepos.NewMockProductRepo(ctrl)

	MockRepo.EXPECT().CheckProductAvailableById(id).Return(false)

	ProductService:=services.NewProductService(MockRepo)
	product ,error :=ProductService.GetProductById(id)
	assert.NotNil(t, product)
	assert.Equal(t, 0,error)
}

func TestAddNewProduct(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	MockRepo :=mock_inmemoryrepos.NewMockProductRepo(ctrl)
	ProductWithoutId:=models.Product{Name: "stump",Description: "wood",Price: 256,Quantity: 1}
	ProductWithId:=models.Product{Id:1,Name: "stump",Description: "wood",Price: 256,Quantity: 1}
	MockRepo.EXPECT().AddProduct(ProductWithoutId).Return(ProductWithId)
	ProductService:=services.NewProductService(MockRepo)
	ReturnedProduct:=ProductService.AddNewProduct(ProductWithoutId)
	assert.NotNil(t, ReturnedProduct)
	assert.Equal(t, 1,ReturnedProduct.Id)
	assert.Equal(t, ProductWithId.Description,ReturnedProduct.Description)
	assert.Equal(t, ProductWithId.Price,ReturnedProduct.Price)
	assert.Equal(t, ProductWithId.Quantity,ReturnedProduct.Quantity)
	assert.Equal(t, ProductWithId.Name,ReturnedProduct.Name)
}

func TestGetAvailableProduct(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	MockRepo :=mock_inmemoryrepos.NewMockProductRepo(ctrl)
	//ProductWithoutId:=models.Product{Name: "stump",Description: "wood",Price: 256,Quantity: 1}
	//ProductWithId:=models.Product{Id:1,Name: "stump",Description: "wood",Price: 256,Quantity: 1}
	MockRepo.EXPECT().GetAvailableProducts().Return(nil)
	ProductService:=services.NewProductService(MockRepo)
	ReturnedProduct:=ProductService.GetAvailableProduct()
	assert.Nil(t, ReturnedProduct)
}

func TestReduceProductQuantity_ProductUnavailable(t *testing.T)  {
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	MockRepo :=mock_inmemoryrepos.NewMockProductRepo(ctrl)
	MockRepo.EXPECT().CheckProductAvailableById(1).Return(false)
	ProductService:=services.NewProductService(MockRepo)
	Response:=ProductService.ReduceProductQuantity(models.Product{},1)
	assert.Equal(t,1 ,Response)
}

func TestReduceProductQuantity_ProductInSufficient(t *testing.T)  {
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	MockRepo :=mock_inmemoryrepos.NewMockProductRepo(ctrl)
	MockRepo.EXPECT().CheckProductAvailableById(1).Return(true)
	MockRepo.EXPECT().GetProductQuantityById(1).Return(10)
	ProductService:=services.NewProductService(MockRepo)
	Response:=ProductService.ReduceProductQuantity(models.Product{Quantity: 20},1)
	assert.Equal(t,2 ,Response)
}

func TestReduceProductQuantity_ProductSufficient(t *testing.T)  {
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	MockRepo :=mock_inmemoryrepos.NewMockProductRepo(ctrl)
	MockRepo.EXPECT().CheckProductAvailableById(1).Return(true)
	MockRepo.EXPECT().GetProductQuantityById(1).Return(25)
	MockRepo.EXPECT().ReduceProductQuantity(models.Product{Quantity: 20},1).Return()
	ProductService:=services.NewProductService(MockRepo)
	Response:=ProductService.ReduceProductQuantity(models.Product{Quantity: 20},1)
	assert.Equal(t,3 ,Response)
}


