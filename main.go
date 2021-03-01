package main

import (
	app "github.com/todo-app-golang/app"
)

func main() {
	a := &app.App{}
	a.InitDb()
	a.InitControllers()
	a.Serve()
}
