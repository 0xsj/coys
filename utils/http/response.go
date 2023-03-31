package http

type Response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

func NewResponse() {

}

func Bytes() {}

func String() {}

func SendResponse() {}

// 200
func StatusOK() {}

// 204
func StatusNoContent() {}

// 400
func StatusBadRequest() {}

// 404
func StatusNotFound() {}

// 405
func StatusMethodNotAllowed() {}

// 409
func StatusConflict() {}

// 500
func StatusInternalServerError() {}
