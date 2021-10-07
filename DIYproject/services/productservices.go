package services

import (
	"awesomeProject1/errors"
	"awesomeProject1/models"
	"awesomeProject1/repos/repointerfaces"
	"time"
)

type ProductService interface {
	GetProductById(id int) (models.Product, error)
	AddNewProduct(product models.Product) models.Product
	GetAllProducts() []models.Product
	GetAvailableProducts() []models.Product
	ReduceProductQuantity(quantity int, id int) error
	IncreaseProductQuantity(quantity int, id int) error
	AddNewSalesRecord(SaleRecord models.SalesRecord)
	GetTop5Product() []models.TopProductResponse
	GetCurrentTime() time.Time
}

type ProductServiceImpl struct {
	Repos repointerfaces.Repo
}

func NewProductService(ProductRepo repointerfaces.Repo) ProductService {
	return &ProductServiceImpl{ProductRepo}
}

func (ps ProductServiceImpl) GetCurrentTime() time.Time {
	return time.Now()
}

func (ps ProductServiceImpl) GetProductById(id int) (models.Product, error) {
	if !ps.Repos.CheckProductAvailableById(id) {
		return models.Product{}, errors.ProductNotFound
	}
	product := ps.Repos.GetProductById(id)
	return product, nil
}

func (ps ProductServiceImpl) AddNewProduct(product models.Product) models.Product {
	return ps.Repos.AddProduct(product)
}

func (ps ProductServiceImpl) AddNewSalesRecord(SaleRecord models.SalesRecord) {
	ps.Repos.AddSalesRecord(SaleRecord)
}

func (ps ProductServiceImpl) GetTop5Product() []models.TopProductResponse {
	return ps.Repos.GetTop5Products()
}

func (ps ProductServiceImpl) GetAllProducts() []models.Product {
	return ps.Repos.GetAllProducts()
}

func (ps ProductServiceImpl) GetAvailableProducts() []models.Product {
	return ps.Repos.GetAvailableProducts()
}

func (ps ProductServiceImpl) ReduceProductQuantity(quantity int, id int) error {
	if !ps.Repos.CheckProductAvailableById(id) {
		return errors.ProductNotFound
	} else if ps.Repos.GetProductQuantityById(id) < quantity {
		return errors.InsufficientProduct
	} else {
		ps.Repos.ReduceProductQuantity(quantity, id)
		return nil
	}
}

func (ps ProductServiceImpl) IncreaseProductQuantity(quantity int, id int) error {
	if !ps.Repos.CheckProductAvailableById(id) {
		return errors.ProductNotFound
	}
	ps.Repos.IncreaseProductQuantity(quantity, id)
	return nil
}
