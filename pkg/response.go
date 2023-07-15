package pkg

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func SendResponse(w http.ResponseWriter, data ApiResponse) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err 		:= encoder.Encode(data)
	PanicIfError(err)
}
