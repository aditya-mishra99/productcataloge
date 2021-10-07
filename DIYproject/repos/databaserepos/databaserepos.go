package databaserepos

import (
	"awesomeProject1/config"
	"awesomeProject1/models"
	"github.com/jinzhu/gorm"
)

type ProductRepo interface {
	GetAvailableProducts() []models.Product
	ReduceProductQuantity(quantity int, id int)
	IncreaseProductQuantity(quantity int, id int)
	AddProduct(product models.Product) models.Product
	GetProductQuantityById(id int) int
	GetAllProducts() []models.Product
	CheckProductAvailableById(id int) bool
	GetProductById(id int) models.Product
	AddSalesRecord(SaleRecord models.SalesRecord)
	GetTop5Products() []models.TopProductResponse
}

type ProductRepoImpl struct {
	Datastore *config.DatabaseDatastore
}

func NewProductRepo(datastore *config.DatabaseDatastore) *ProductRepoImpl {
	return &ProductRepoImpl{datastore}
}

func (pr ProductRepoImpl) GetProductById(id int) models.Product {
	var p models.Product
	pr.Datastore.Products.First(&p, id)
	return p
}

func (pr ProductRepoImpl) CheckProductAvailableById(id int) bool {

	if err := pr.Datastore.Products.Where("Id=?", id).Find(&models.Product{}).Error; gorm.IsRecordNotFoundError(err) {
		return false
	}
	return true
}

func (pr ProductRepoImpl) GetAllProducts() []models.Product {
	AllProduct := make([]models.Product, 0)
	pr.Datastore.Products.Find(&AllProduct)
	return AllProduct
}

func (pr ProductRepoImpl) DeleteProduct(id int) {
	var product models.Product
	pr.Datastore.Products.First(&product, id)
	pr.Datastore.Products.Delete(&product)

}

func (pr ProductRepoImpl) GetProductQuantityById(id int) int {
	var p models.Product
	pr.Datastore.Products.First(&p, id)
	return p.Quantity
}

func (pr ProductRepoImpl) AddProduct(product models.Product) models.Product {
	pr.Datastore.Products.NewRecord(product)
	pr.Datastore.Products.Create(&product)
	return product
}

func (pr ProductRepoImpl) AddSalesRecord(SaleRecord models.SalesRecord) {
	pr.Datastore.Products.NewRecord(SaleRecord)
	pr.Datastore.Products.Create(&SaleRecord)
}

func (pr ProductRepoImpl) GetTop5Products() []models.TopProductResponse {
	var TopProduct []models.TopProductResponse
	sqlStr := "SELECT product_id,SUM(quantity_sold) AS Total FROM (SELECT product_id,quantity_sold FROM sales_records WHERE sales_time> NOW() - INTERVAL '1 hour') Temp GROUP BY product_id ORDER BY Total DESC LIMIT 5"
	rows, err := pr.Datastore.Products.Raw(sqlStr).Rows()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var p models.TopProductResponse
		var Id int
		var QuantitySold int
		err = rows.Scan(&Id, &QuantitySold)
		p.ProductId = Id
		p.QuantitySold = QuantitySold
		TopProduct = append(TopProduct, p)
	}
	return TopProduct
}

func (pr ProductRepoImpl) IncreaseProductQuantity(quantity int, id int) {
	Product := pr.GetProductById(id)
	Product.Quantity += quantity
	pr.Datastore.Products.Save(Product)
}

func (pr ProductRepoImpl) ReduceProductQuantity(quantity int, id int) {
	Product := pr.GetProductById(id)
	Product.Quantity -= quantity
	pr.Datastore.Products.Save(Product)
}

func (pr ProductRepoImpl) GetAvailableProducts() []models.Product {
	ProductArray := make([]models.Product, 0)
	for _, Product := range pr.GetAllProducts() {
		if Product.Quantity > 0 {
			ProductArray = append(ProductArray, Product)
		}
	}
	return ProductArray
}
