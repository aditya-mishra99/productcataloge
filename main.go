package main

import (
	"awesomeProject1/controller"
	"awesomeProject1/models"
	"awesomeProject1/repos/inmemoryrepos"
	"awesomeProject1/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)



func main() {

	r := mux.NewRouter()
    InmemoryDatastore := models.InitialiseInmemoryDatastore()
	ProductRepo :=inmemoryrepos.NewProductRepo(InmemoryDatastore)
	ProductService := services.NewProductService(ProductRepo)
	ProductController := controller.Initialise(ProductService)

	r.HandleFunc("/products", ProductController.GetAllProduct).Methods("GET")
	r.HandleFunc("/products/{id}", ProductController.GetProductById).Methods("GET")
	r.HandleFunc("/availableProducts", ProductController.GetAvailableProduct).Methods("GET")
	r.HandleFunc("/products", ProductController.CreateProduct).Methods("PUT")
	r.HandleFunc("/products/purchase/{id}", ProductController.BuyProduct).Methods("PUT")
	r.HandleFunc("/products/update/{id}", ProductController.IncreaseQuantity).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))

}
