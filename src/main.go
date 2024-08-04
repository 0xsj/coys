package main

import "github.com/0xsj/coys/src/libs/app"

func main() {
	module := NewAppModule()
	app := app.NewApp(module)
	port := 8080
	app.Start(port)
}