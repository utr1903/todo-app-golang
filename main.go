package main

import (
	app "github.com/todo-app-golang/app"
)

func main() {
	app := &app.App{}
	app.InitDb()
	app.InitControllers()
	app.Serve()
}
