// injectable.go
package di

import (
	"reflect"
)

// Injectable marks a type as injectable.
type Injectable interface {
	Provide() interface{}
}

// RegisterInjectable registers an injectable type with the container.
func RegisterInjectable(instance Injectable) {
	GlobalContainer.RegisterProvider(reflect.TypeOf(instance).String(), instance.Provide())
}

