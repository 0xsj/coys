// errors.go
package config

import "fmt"

var (
	ErrConfigFileNotFound = fmt.Errorf("config file not found")
	ErrConfigLoadFailed   = fmt.Errorf("failed to load config")
)
