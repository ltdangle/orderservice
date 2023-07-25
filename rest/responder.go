package rest

import (
	"encoding/json"
	"net/http"
	"time"
)

// JSON response object.
type Response struct {
	Date    string      `json:"date"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

func NewResponse() *Response {
	return &Response{}
}

// Responder.
type Responder struct {
	dateFormat string
}

func NewResponder(dateFormat string) *Responder {
	return &Responder{dateFormat: dateFormat}
}

func (rspndr *Responder) Error(w http.ResponseWriter, statusCode int, message string) {
	// Build json response.
	r := NewResponse()
	r.Date = time.Now().Format(rspndr.dateFormat)
	r.Message = message

	// Send response.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(r)
}

func (rspndr *Responder) Success(w http.ResponseWriter, payload interface{}) {
	// Build json response.
	r := NewResponse()
	r.Date = time.Now().Format(rspndr.dateFormat)
	r.Message = "ok"
	r.Payload = payload

	// Send response.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(r)
}
