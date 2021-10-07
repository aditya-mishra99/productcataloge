package utility

import (
	"awesomeProject1/models"
	"encoding/json"
	"net/http"
)

func ResponseWithMessageStatus(w http.ResponseWriter,status int,message string){
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"message":message,
	})
}




func SetJsonContentType(w http.ResponseWriter){
	w.Header().Set("Content-Type", "application/json")
}




func ResponseWithProductArray(w http.ResponseWriter,p []models.Product,status int){
	json.NewEncoder(w).Encode(p)

}

func ResponseWithTopProductArray(w http.ResponseWriter,p []models.TopProductResponse,status int){
	json.NewEncoder(w).Encode(p)
	w.WriteHeader(status)
}

func ResponseWithProduct(w http.ResponseWriter,p models.Product,status int){
	json.NewEncoder(w).Encode(p)
	w.WriteHeader(status)
}


