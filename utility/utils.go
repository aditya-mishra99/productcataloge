package utility

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter,status int,message string){
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"message":message,
	})
}