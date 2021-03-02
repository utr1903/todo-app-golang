package app

import (
	"database/sql"
	"log"
	"net/http"

	// mysql import
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/todo-app-golang/controllers"
	"github.com/todo-app-golang/controllers/todoitemmodule"
	"github.com/todo-app-golang/controllers/todolistmodule"
	"github.com/todo-app-golang/controllers/usersmodule"
)

// App : DB and Controllers
type App struct {
	Db     *sql.DB
	Router *mux.Router
}

// InitDb : Initializes the Db connection
func (a *App) InitDb() {
	db, err := sql.Open("mysql", "utr1903:admin@(127.0.0.1:3306)/Todo?parseTime=true")
	a.Db = db
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}

// InitControllers : Initializes the controllers
func (a *App) InitControllers() {
	r := mux.NewRouter()
	a.Router = r

	initUserController(a)
	initTodoListController(a)
	initTodoItemController(a)
}

func initUserController(a *App) {
	b := &controllers.Controller{Db: a.Db}
	c := &usersmodule.UsersController{Base: b}

	a.Router.HandleFunc("/users/get", c.GetUsers).Methods("GET")
	a.Router.HandleFunc("/users/create", c.CreateUser).Methods("POST")
	a.Router.HandleFunc("/user/update", c.GetUser).Methods("POST")
}

func initTodoListController(a *App) {
	b := &controllers.Controller{Db: a.Db}
	c := &todolistmodule.TodoListController{Base: b}

	a.Router.HandleFunc("/lists/get", c.GetLists).Methods("GET")
	a.Router.HandleFunc("/lists/create", c.CreateTodoList).Methods("POST")
	a.Router.HandleFunc("/list/get", c.GetList).Methods("POST")
	a.Router.HandleFunc("/list/update", c.UpdateTodoList).Methods("POST")
}

func initTodoItemController(a *App) {
	b := &controllers.Controller{Db: a.Db}
	c := &todoitemmodule.TodoItemController{Base: b}

	a.Router.HandleFunc("/items/get", c.GetTodoItems).Methods("GET")
	a.Router.HandleFunc("/items/create", c.CreateTodoItem).Methods("POST")
	a.Router.HandleFunc("/item/get", c.GetTodoItem).Methods("POST")
	a.Router.HandleFunc("/item/update", c.UpdateTodoItem).Methods("POST")
}

// Serve : Runs web server
func (a *App) Serve() {
	http.ListenAndServe(":8080", a.Router)
}
