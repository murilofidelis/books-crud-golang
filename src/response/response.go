package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func Json(w http.ResponseWriter, statuCode int, data interface{}) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statuCode)

	log.Print(data)
	if data == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	Json(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: err.Error(),
	})
}
