package main

import (
	"github.com/0xsj/coys/src/libs/app"
)

func main() {
	appModule := NewAppModule()
	app := app.NewApp(appModule)
	port := 8080
	app.Start(port)
}