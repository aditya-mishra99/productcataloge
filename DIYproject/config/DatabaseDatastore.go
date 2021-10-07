package config

import (
	"awesomeProject1/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type DatabaseDatastore struct {
	Products *gorm.DB //make db
}

func InitialiseDatabaseDatastore() *DatabaseDatastore {
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=ProductCatalog sslmode=disable password=Dunzo@123 port=5432")
	db1, err := gorm.Open("postgres", dbURI)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database successfully")
	}
	db1.AutoMigrate(&models.Product{})
	db1.AutoMigrate(&models.SalesRecord{})
	return &DatabaseDatastore{db1}
}
