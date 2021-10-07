package repointerfaces

import "awesomeProject1/models"

type Repo interface {
	GetProductById(id int) models.Product
	CheckProductAvailableById(id int) bool
	GetAllProducts() []models.Product
	GetProductQuantityById(id int) int
	AddProduct(product models.Product) models.Product
	IncreaseProductQuantity(quantity int, id int)
	ReduceProductQuantity(quantity int, id int)
	GetAvailableProducts() []models.Product
	GetTop5Products() []models.TopProductResponse
	AddSalesRecord(SaleRecord models.SalesRecord)
}
