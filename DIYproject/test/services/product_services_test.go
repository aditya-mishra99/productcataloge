package services

import (
	"awesomeProject1/errors"
	"awesomeProject1/models"
	"awesomeProject1/services"
	mock_databaserepos "awesomeProject1/test/mock/repos"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProductById_AvailableProduct(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	id:=1
	MockRepo:= mock_databaserepos.NewMockProductRepo(ctrl)
	MockRepo.EXPECT().GetProductById(id).Return(models.Product{
		Id : 1,
		Price: 22,
		Name : "Balls",
		Quantity: 10,
		Description: "adidas",
	})
	MockRepo.EXPECT().CheckProductAvailableById(id).Return(true)
	ProductService:= services.NewProductService(MockRepo)
	product ,error :=ProductService.GetProductById(id)
	assert.NotNil(t, product)
	assert.Equal(t,nil,error)
}

func TestGetProductById_UnAvailableProduct(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	id:=1
	MockRepo := mock_databaserepos.NewMockProductRepo(ctrl)

	MockRepo.EXPECT().CheckProductAvailableById(id).Return(false)

	ProductService:= services.NewProductService(MockRepo)
	product ,error :=ProductService.GetProductById(id)
	assert.Equal(t, 0,product.Id)
	assert.Equal(t, 0,product.Price)
	assert.Equal(t, "",product.Name)
	assert.Equal(t, 0,product.Quantity)
	assert.Equal(t, "",product.Description)
	assert.Equal(t, errors.ProductNotFound,error)

}

func TestAddNewProduct(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	MockRepo := mock_databaserepos.NewMockProductRepo(ctrl)
	ProductWithoutId:=models.Product{Name: "stump",Description: "wood",Price: 256,Quantity: 1}
	ProductWithId:=models.Product{Id:1,Name: "stump",Description: "wood",Price: 256,Quantity: 1}
	MockRepo.EXPECT().AddProduct(ProductWithoutId).Return(ProductWithId)
	ProductService:= services.NewProductService(MockRepo)
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
	MockRepo := mock_databaserepos.NewMockProductRepo(ctrl)
	//ProductWithoutId:=models.Product{Name: "stump",Description: "wood",Price: 256,Quantity: 1}
	//ProductWithId:=models.Product{Id:1,Name: "stump",Description: "wood",Price: 256,Quantity: 1}
	MockRepo.EXPECT().GetAvailableProducts().Return(nil)
	ProductService:= services.NewProductService(MockRepo)
	ReturnedProduct:=ProductService.GetAvailableProducts()
	assert.Nil(t, ReturnedProduct)
}

func TestReduceProductQuantity_ProductUnavailable(t *testing.T)  {
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	MockRepo := mock_databaserepos.NewMockProductRepo(ctrl)
	MockRepo.EXPECT().CheckProductAvailableById(1).Return(false)
	ProductService:= services.NewProductService(MockRepo)
	Response:=ProductService.ReduceProductQuantity(0,1)
	assert.Equal(t,errors.ProductNotFound ,Response)
}

func TestReduceProductQuantity_ProductInSufficient(t *testing.T)  {
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	MockRepo := mock_databaserepos.NewMockProductRepo(ctrl)
	MockRepo.EXPECT().CheckProductAvailableById(1).Return(true)
	MockRepo.EXPECT().GetProductQuantityById(1).Return(10)
	ProductService:= services.NewProductService(MockRepo)
	Response:=ProductService.ReduceProductQuantity(20,1)
	assert.Equal(t,errors.InsufficientProduct ,Response)
}

func TestReduceProductQuantity_ProductSufficient(t *testing.T)  {
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	MockRepo := mock_databaserepos.NewMockProductRepo(ctrl)
	MockRepo.EXPECT().CheckProductAvailableById(1).Return(true)
	MockRepo.EXPECT().GetProductQuantityById(1).Return(25)
	MockRepo.EXPECT().ReduceProductQuantity(20,1).Return()
	ProductService:= services.NewProductService(MockRepo)
	Response:=ProductService.ReduceProductQuantity(20,1)
	assert.Equal(t,nil ,Response)
}


func TestIncreaseProductQuantity_ProductUnavailable(t *testing.T)  {
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	MockRepo := mock_databaserepos.NewMockProductRepo(ctrl)
	MockRepo.EXPECT().CheckProductAvailableById(1).Return(false)
	ProductService:= services.NewProductService(MockRepo)
	Response:=ProductService.ReduceProductQuantity(0,1)
	assert.Equal(t,errors.ProductNotFound ,Response)
}

func TestIncreaseProductQuantity_ProductAvailable(t *testing.T)  {
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	MockRepo := mock_databaserepos.NewMockProductRepo(ctrl)
	MockRepo.EXPECT().CheckProductAvailableById(2).Return(true)
	MockRepo.EXPECT().IncreaseProductQuantity(1,2).Return()
	ProductService:= services.NewProductService(MockRepo)
	Response:=ProductService.IncreaseProductQuantity(1,2)
	assert.Equal(t,nil ,Response)
}
