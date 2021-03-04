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

	a.Router.HandleFunc("/users/get", c.GetUsers).Methods("GET", "OPTIONS")
	a.Router.HandleFunc("/users/create", c.CreateUser).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/user/get", c.GetUser).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/user/update", c.UpdateUser).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/user/delete", c.DeleteUser).Methods("POST", "OPTIONS")
}

func initTodoListController(a *App) {
	b := &controllers.Controller{Db: a.Db}
	c := &todolistmodule.TodoListController{Base: b}

	a.Router.HandleFunc("/lists/get", c.GetLists).Methods("GET", "OPTIONS")
	a.Router.HandleFunc("/lists/create", c.CreateTodoList).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/list/get", c.GetList).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/list/update", c.UpdateTodoList).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/list/delete", c.DeleteTodoList).Methods("POST", "OPTIONS")
}

func initTodoItemController(a *App) {
	b := &controllers.Controller{Db: a.Db}
	c := &todoitemmodule.TodoItemController{Base: b}

	a.Router.HandleFunc("/items/get", c.GetTodoItems).Methods("GET", "OPTIONS")
	a.Router.HandleFunc("/items/create", c.CreateTodoItem).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/item/get", c.GetTodoItem).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/item/update", c.UpdateTodoItem).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/item/delete", c.DeleteTodoItem).Methods("POST", "OPTIONS")
}

// RouterWithCORS : To prevent getting CORS errors from Angular UI
type RouterWithCORS struct {
	r *mux.Router
}

// Serve : Runs web server
func (a *App) Serve() {
	http.ListenAndServe(":8080", &RouterWithCORS{a.Router})
}

// ServeHTTP : A middleware to add necessary headers in order not to get CORS error
func (s *RouterWithCORS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		// w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	// Stop here for a Preflighted OPTIONS request.
	if r.Method == "OPTIONS" {
		return
	}

	// Lets Gorilla work
	s.r.ServeHTTP(w, r)
}
