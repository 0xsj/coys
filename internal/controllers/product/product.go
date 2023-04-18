package product

import "github.com/sjtommylee/go-dynamodb/internal/repository/adapter"

type Controller struct {
	repository adapter.Interface
}

type Interface interface {
	ListOne()
	ListAll()
	Create()
	Update()
	Remove()
}

func NewController() {}

func ListOne() {}

func ListAll() {}

func Create() {}

func Update() {}

func Remove() {}
