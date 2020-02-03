package main

import (
	"context"
	"encoding/json"
	"net/http"

	"code_study_microservice/http-server-client/entities"
)

type validationContextKey string

type requestHandler struct {
	next http.Handler
}

func newRequestHandler(next http.Handler) http.Handler {
	return requestHandler{next}
}

func (h requestHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request entities.RequestData
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	r = r.WithContext(ctx)

	h.next.ServeHTTP(rw, r)
}
