package main

import (
	"encoding/json"
	"net/http"
)

type responseData struct{
	Message string
}

type responseHandler struct{
}

func newResponseHandler() http.Handler {
	return responseHandler{}
}

func (h responseHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(validationContextKey("name")).(string)
	response := responseData{Message:"Hello" + name}

	encoder := json.NewEncoder(rw)
	err := encoder.Encode(&response)
	if err != nil {
		panic("Ooops")
	}
}
