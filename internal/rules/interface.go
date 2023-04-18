package rules

type Interface interface {
	ConvertIoRtoStruct()
	GetMock() interface{}
	Migrate()
	Validate()
}
