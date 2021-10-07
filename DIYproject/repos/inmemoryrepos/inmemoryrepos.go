package inmemoryrepos

import (
	"awesomeProject1/config"
	"awesomeProject1/models"
	"sync"
)

type ProductRepo interface {
	GetAvailableProducts() []models.Product
	ReduceProductQuantity(quantity int, id int)
	IncreaseProductQuantity(quantity int, id int)
	AddProduct(product models.Product) models.Product
	GetProductQuantityById(id int) int
	GetAllProduct() []models.Product
	CheckProductAvailableById(id int) bool
	GetProductById(id int) models.Product
}

type ProductRepoImpl struct {
	Datastore *config.InmemoryDatastore
	mutex     *sync.Mutex
}

func NewProductRepo(datastore *config.InmemoryDatastore) *ProductRepoImpl {
	return &ProductRepoImpl{datastore, &sync.Mutex{}}
}

func (pr ProductRepoImpl) GetProductById(id int) models.Product {
	pr.mutex.Lock()
	p := pr.Datastore.Products[id]
	pr.mutex.Unlock()
	return p
}

func (pr ProductRepoImpl) CheckProductAvailableById(id int) bool {
	pr.mutex.Lock()
	if _, ok := pr.Datastore.Products[id]; !ok {
		pr.mutex.Unlock()
		return false
	}
	pr.mutex.Unlock()
	return true
}

func (pr ProductRepoImpl) GetAllProducts() []models.Product {
	TempProduct := make([]models.Product, 0)
	pr.mutex.Lock()
	for _, val := range pr.Datastore.Products {
		TempProduct = append(TempProduct, val)
	}
	pr.mutex.Unlock()
	return TempProduct
}

func (pr ProductRepoImpl) GetProductQuantityById(id int) int {
	pr.mutex.Lock()
	quantity := pr.Datastore.Products[id].Quantity
	pr.mutex.Unlock()
	return quantity
}

func (pr ProductRepoImpl) AddProduct(product models.Product) models.Product {
	pr.mutex.Lock()
	id := len(pr.Datastore.Products) + 1
	product.Id = id
	pr.Datastore.Products[id] = product
	pr.mutex.Unlock()
	return product
}

func (pr ProductRepoImpl) IncreaseProductQuantity(quantity int, id int) {
	pr.mutex.Lock()
	ProductNewState := pr.Datastore.Products[id]
	ProductNewState.Quantity += quantity
	pr.Datastore.Products[id] = ProductNewState
	pr.mutex.Unlock()
}

func (pr ProductRepoImpl) ReduceProductQuantity(quantity int, id int) {
	pr.mutex.Lock()
	ProductNewState := pr.Datastore.Products[id]
	ProductNewState.Quantity -= quantity
	pr.Datastore.Products[id] = ProductNewState
	pr.mutex.Unlock()
}

func (pr ProductRepoImpl) GetAvailableProducts() []models.Product {
	ProductArray := make([]models.Product, 0)
	pr.mutex.Lock()
	for _, item := range pr.Datastore.Products {
		if item.Quantity > 0 {
			ProductArray = append(ProductArray, item)
		}
	}
	pr.mutex.Unlock()
	return ProductArray
}

func (pr ProductRepoImpl) GetTop5Products() []models.TopProductResponse {
	//dummy  did for interface uses
	p := make([]models.TopProductResponse, 0)
	return p
}

func (pr ProductRepoImpl) AddSalesRecord(SaleRecord models.SalesRecord) {
	pr.Datastore.Sales = append(pr.Datastore.Sales, SaleRecord)
}
