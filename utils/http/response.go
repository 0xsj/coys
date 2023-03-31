package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

func NewResponse(data interface{}, status int) *Response {
	return &Response{
		Status: status,
		Result: data,
	}
}

func (response *Response) Bytes() []byte {
	data, _ := json.Marshal(response)
	return data
}

func (response *Response) String() string {
	return string(response.Bytes())
}

func (response *Response) SendResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(response.Status)
	_, _ = w.Write(response.Bytes())
	log.Println(response.String())
}

// 200
func StatusOK(w http.ResponseWriter, r *http.Request, data interface{}) {
	response := NewResponse(data, http.StatusOK)
	response.SendResponse(w, r)
}

// 204
func StatusNoContent(w http.ResponseWriter, r *http.Request) {
	response := NewResponse(nil, http.StatusNoContent)
	response.SendResponse(w, r)
}

// 400
func StatusBadRequest(w http.ResponseWriter, r *http.Request, err error) {}

// 404
func StatusNotFound(w http.ResponseWriter) {}

// 405
func StatusMethodNotAllowed(w http.ResponseWriter) {}

// 409
func StatusConflict(w http.ResponseWriter) {}

// 500
func StatusInternalServerError(w http.ResponseWriter) {}
