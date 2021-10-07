package config

import (
	"awesomeProject1/models"
	_ "github.com/lib/pq"
)

type InmemoryDatastore struct{
	Products map[int]models.Product
	Sales    []models.SalesRecord
}

func InitialiseInmemoryDatastore() *InmemoryDatastore {
	products:=make(map[int]models.Product)
	sales:=make([]models.SalesRecord,0)
	products[1]= models.Product{Id: 1, Name: "bat", Description: "mrf bat", Price: 220, Quantity: 10}
	products[2]= models.Product{Id: 2, Name: "stump", Description: "mrf stump", Price: 22, Quantity: 10}
	products[3]= models.Product{Id: 3, Name: "helmet", Description: "solid helmet", Price: 220, Quantity: 10}
	p:= InmemoryDatastore{products,sales}
	return &p
}

//type InmemoryDatastore struct{
//	Products *gorm.DB
//}
//
//func InitialiseInmemoryDatastore() *InmemoryDatastore {
//	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=ProductCatalog sslmode=disable password=Dunzo@123 port=5432")
//	db1 ,err :=gorm.Open("postgres", dbURI)
//
//	if err != nil {
//		panic(err)
//	} else {
//		fmt.Println("Connected to database successfully")
//	}
//
//	db1.AutoMigrate(&Product{})
//	db1.AutoMigrate(&SalesRecord{})
//	return  &InmemoryDatastore{db1}
//}

