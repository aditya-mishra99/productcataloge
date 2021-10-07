package main

import (
	"awesomeProject1/config"
	"awesomeProject1/controller"
	"awesomeProject1/repos/databaserepos"
	"awesomeProject1/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)



func main() {

	r := mux.NewRouter()

	//running inmemory apis

	//InmemoryDatastore := config.InitialiseInmemoryDatastore()
	//ProductRepo:=InmemoryRepos.NewProductRepo(InmemoryDatastore)


    // running database apis

    DatabaseDatastore := config.InitialiseDatabaseDatastore()
	ProductRepo := databaserepos.NewProductRepo(DatabaseDatastore)

	ProductService := services.NewProductService(ProductRepo)
	ProductController := controller.Initialise(ProductService)


	r.HandleFunc("/products", ProductController.GetAllProducts).Methods("GET")
	r.HandleFunc("/products/topsold", ProductController.GetTop5Product).Methods("GET")
	r.HandleFunc("/products/{id}", ProductController.GetProductById).Methods("GET")
	r.HandleFunc("/availableProducts", ProductController.GetAvailableProducts).Methods("GET")
	r.HandleFunc("/products", ProductController.CreateProduct).Methods("POST")
	r.HandleFunc("/products/purchase/{id}", ProductController.BuyProduct).Methods("PUT")
	r.HandleFunc("/products/update/{id}", ProductController.IncreaseQuantity).Methods("PUT")


	log.Fatal(http.ListenAndServe(":8000", r))

}
