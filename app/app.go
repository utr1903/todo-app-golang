package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// mysql import
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/todo-app-golang/services"
)

// App : DB and Controllers
type App struct {
	Db     *sql.DB
	Router *mux.Router
}

// Hi : For test purposes
func (a *App) Hi() {
	fmt.Println("Hi")
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
	r.HandleFunc("/users", a.getUsers).Methods("GET")
	r.HandleFunc("/lists", a.getLists).Methods("GET")
	a.Router = r
}

func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {
	us := &services.UserService{}
	users, err := us.GetUsers(a.Db)

	if err != nil {
		createResponse(w, http.StatusBadRequest, nil)
	}

	createResponse(w, http.StatusOK, users)
}

func (a *App) getLists(w http.ResponseWriter, r *http.Request) {
	tls := &services.TodoListService{}
	lists, err := tls.GetLists(a.Db)

	if err != nil {
		createResponse(w, http.StatusBadRequest, nil)
	}

	createResponse(w, http.StatusOK, lists)
}

func createResponse(w http.ResponseWriter, code int, dto interface{}) {
	response, _ := json.Marshal(dto)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Serve : Runs web server
func (a *App) Serve() {
	http.ListenAndServe(":8080", a.Router)
}
