package app

import "net/http"

type Module interface {
	RegisterRoutes(router *http.ServeMux) 
}


