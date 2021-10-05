package services

import (
	"awesomeProject1/models"
	"awesomeProject1/repos/inmemoryrepos"
)

type ProductService interface {
	GetProductById(id int) (p models.Product, err int)
	AddNewProduct(product models.Product) models.Product
	GetAllProduct() []models.Product
	GetAvailableProduct() []models.Product
	ReduceProductQuantity(quantity models.Product,id int) int
	IncreaseProductQuantity(quantity models.Product,id int)
}


type ProductServiceImpl struct{
	Repos inmemoryrepos.ProductRepo
}

func NewProductService(ProductRepo inmemoryrepos.ProductRepo) ProductService {
	return &ProductServiceImpl{ProductRepo}
}


func (ps ProductServiceImpl) GetProductById(id int) (p models.Product, err int) {
	if !ps.Repos.CheckProductAvailableById(id){
		return models.Product{} , 0
	}
	product:=ps.Repos.GetProductById(id)
	return product, 1
}



func (ps ProductServiceImpl) AddNewProduct(product models.Product) models.Product {
	return ps.Repos.AddProduct(product)
}



func (ps ProductServiceImpl) GetAllProduct() []models.Product {
	return ps.Repos.GetAllProduct()
}


func (ps ProductServiceImpl) GetAvailableProduct() []models.Product {
	return ps.Repos.GetAvailableProducts()
}

func (ps ProductServiceImpl) ReduceProductQuantity(quantity models.Product,id int) int {
	if !ps.Repos.CheckProductAvailableById(id){
		return 1
	} else if ps.Repos.GetProductQuantityById(id)<quantity.Quantity {
		return 2
	}else {
		ps.Repos.ReduceProductQuantity(quantity,id)
		return 3
	}
}


func (ps ProductServiceImpl) IncreaseProductQuantity(quantity models.Product,id int) {
	 ps.Repos.IncreaseProductQuantity(quantity,id)
}


