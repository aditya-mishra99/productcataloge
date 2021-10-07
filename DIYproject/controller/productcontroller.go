package controller

import (
	"awesomeProject1/errors"
	"awesomeProject1/models"
	"awesomeProject1/services"
	"awesomeProject1/utility"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ProductController struct { ///change name
	ProductService services.ProductService // change name
}

func Initialise(p services.ProductService) *ProductController {
	return &ProductController{p}
}

// add try catch and corres status in all

func (pc ProductController) GetProductById(w http.ResponseWriter, r *http.Request) {
	utility.SetJsonContentType(w)
	params := mux.Vars(r)
	id, error1 := strconv.Atoi(params["id"])
	if error1 != nil {
		utility.ResponseWithMessageStatus(w, 400, "Invalid id type")
		return
	}
	IdProduct, error2 := pc.ProductService.GetProductById(id) // change err
	if error2 != nil {
		utility.ResponseWithMessageStatus(w, http.StatusOK, "Product Not available")
		return
	}
	utility.ResponseWithProduct(w, IdProduct, http.StatusOK)
}

func (pc ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	utility.SetJsonContentType(w)
	var product models.Product
	error1 := json.NewDecoder(r.Body).Decode(&product)
	if error1 != nil {
		utility.ResponseWithMessageStatus(w, http.StatusBadRequest, "invalid json was passed")
		return
	}
	p := pc.ProductService.AddNewProduct(product)
	utility.ResponseWithProduct(w, p, http.StatusOK)
}

func (pc ProductController) IncreaseQuantity(w http.ResponseWriter, r *http.Request) {
	utility.SetJsonContentType(w)
	var product models.Product
	error1 := json.NewDecoder(r.Body).Decode(&product)
	if error1 != nil {
		utility.ResponseWithMessageStatus(w, http.StatusBadRequest, "invalid json was passed")
		return
	}
	params := mux.Vars(r)
	id, error2 := strconv.Atoi(params["id"])
	if error2 != nil {
		utility.ResponseWithMessageStatus(w, http.StatusBadRequest, "invalid id was passed")
		return
	}
	error3 := pc.ProductService.IncreaseProductQuantity(product.Quantity, id)
	if error3 != nil {
		utility.ResponseWithMessageStatus(w, http.StatusBadRequest, "product does not exist")
		return
	}
	utility.ResponseWithMessageStatus(w, http.StatusOK, "product update successful")
}

func (pc ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	utility.SetJsonContentType(w)
	p := pc.ProductService.GetAllProducts()
	utility.ResponseWithProductArray(w, p, http.StatusOK)
}

func (pc ProductController) GetAvailableProducts(w http.ResponseWriter, r *http.Request) {
	utility.SetJsonContentType(w)
	p := pc.ProductService.GetAvailableProducts()
	utility.ResponseWithProductArray(w, p, http.StatusOK)
}

func (pc ProductController) BuyProduct(w http.ResponseWriter, r *http.Request) {
	utility.SetJsonContentType(w)
	params := mux.Vars(r)
	var value models.Product
	error1 := json.NewDecoder(r.Body).Decode(&value)
	if error1 != nil {
		utility.ResponseWithMessageStatus(w, 400, "invalid json was passed")
		return
	}
	id, _ := strconv.Atoi(params["id"])
	error := pc.ProductService.ReduceProductQuantity(value.Quantity, id)
	if error == errors.ProductNotFound {
		utility.ResponseWithMessageStatus(w, http.StatusOK, "product not available")
		return
	} else if error == errors.InsufficientProduct {
		utility.ResponseWithMessageStatus(w, http.StatusOK, "sufficient quantity not available")
		return
	} else {
		utility.ResponseWithMessageStatus(w, http.StatusOK, "product purchase successful")
		PurchaseTime := pc.ProductService.GetCurrentTime()
		pc.ProductService.AddNewSalesRecord(models.SalesRecord{ProductId: id, QuantitySold: value.Quantity, SalesTime: PurchaseTime})
		return
	}
}

func (pc ProductController) GetTop5Product(w http.ResponseWriter, r *http.Request) {
	utility.SetJsonContentType(w)
	Top5Products := pc.ProductService.GetTop5Product()
	length := len(Top5Products)
	if length == 0 {
		utility.ResponseWithMessageStatus(w, http.StatusOK, "no purchase in last one hour")
		return
	}
	utility.ResponseWithTopProductArray(w, Top5Products, http.StatusOK)
}
