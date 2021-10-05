package controller

import (
	"awesomeProject1/models"
	"awesomeProject1/services"
	"awesomeProject1/utility"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ProductController interface{
	GetProductById(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	IncreaseQuantity(w http.ResponseWriter, r *http.Request)
	GetAllProduct(w http.ResponseWriter, r *http.Request)
	GetAvailableProduct(w http.ResponseWriter, r *http.Request)
	BuyProduct(w http.ResponseWriter, r *http.Request)
}



type ProductControllerImpl struct{
	 ps  services.ProductService
}

func Initialise(p services.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{p}
}

func (pc ProductControllerImpl) GetProductById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
	    id, _ := strconv.Atoi(params["id"])
		IdProduct ,err :=pc.ps.GetProductById(id)
		if err==0 {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]string{
				"message":"Product Not available",
			})
			return
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(IdProduct)
}


func (pc ProductControllerImpl) CreateProduct(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var product models.Product
		_ = json.NewDecoder(r.Body).Decode(&product)
	    p:=pc.ps.AddNewProduct(product)
	    json.NewEncoder(w).Encode(p)
	    w.WriteHeader(200)
}

func (pc ProductControllerImpl) IncreaseQuantity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	pc.ps.IncreaseProductQuantity(product,id)
	w.WriteHeader(200)
	utility.Response(w,200,"Product Update Successful")
}


func (pc ProductControllerImpl) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p:= pc.ps.GetAllProduct()
	json.NewEncoder(w).Encode(p)
	w.WriteHeader(200)
}

func (pc ProductControllerImpl) GetAvailableProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p:=pc.ps.GetAvailableProduct()
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(p)
}

func (pc ProductControllerImpl) BuyProduct(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var value models.Product
	_ = json.NewDecoder(r.Body).Decode(&value)
	id, _ := strconv.Atoi(params["id"])
	state:=pc.ps.ReduceProductQuantity(value,id)
	if state== 1{
		utility.Response(w,400,"Product Not available")
	}else if state == 2 {
		utility.Response(w,400,"sufficient quantity not available")
	}else{
		utility.Response(w,200,"Product purchase successful")
	}
}
