package web

import (
	"encoding/json"
	"net/http"
)

type Responder interface {
	Send(w http.ResponseWriter) error
}
type ResponseAPI struct {
	Status int         `json:"status,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

type ResponseError struct {
	Status int         `json:"status,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

func Success(result interface{}, status int) Responder {
	return &ResponseAPI{
		Status: status,
		Result: result,
	}
}

func Failure(err error, status int) Responder {
	return &ResponseError{
		Status: status,
		Result: err.Error(),
	}
}

func (r *ResponseAPI) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	return json.NewEncoder(w).Encode(r.Result)
}

func (r *ResponseError) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	return json.NewEncoder(w).Encode(r)
}
