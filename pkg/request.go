package pkg

import (
	"encoding/json"
	"net/http"
)

func RequestParse(r *http.Request, body interface{}) {
	decoder := json.NewDecoder(r.Body)
	err 		:= decoder.Decode(body)
	PanicIfError(err)
}