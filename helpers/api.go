package helpers

import (
	"encoding/json"
	"net/http"
)

type BaseResponseBody struct {
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
}

func Respond(w http.ResponseWriter, data BaseResponseBody) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(data.Status)
	json.NewEncoder(w).Encode(data)
}
