package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "utr1903:admin@(127.0.0.1:3306)/Todo?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{ // Query a single user
		var (
			id       string
			username string
			password string
		)

		query := "SELECT id, username, password FROM users WHERE id = ?"
		if err := db.QueryRow(query, "7054d86a-79b8-11eb-b4bd-00090ffe0001").Scan(&id, &username, &password); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password)
	}

	// Routing
	// r := mux.NewRouter()

	// r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	title := vars["title"]
	// 	page := vars["page"]

	// 	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	// })

	// http.ListenAndServe(":80", r)
}
