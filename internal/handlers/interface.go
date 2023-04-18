package handlers

type Interface interface {
	Get()
	Post()
	Put()
	Delete()
	Options()
}
