package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}


func WriteToResponseBody(w http.ResponseWriter, r interface{})  {
	w.Header().Add("Content-Type", "Application/json")
	encoder := json.NewEncoder(w)
    err := encoder.Encode(r)
	PanicIfError(err)
}