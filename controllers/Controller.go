package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// Controller : Base class for controllers
type Controller struct {
	Db *sql.DB
}

// CreateResponse : Generic way for all controllers to create JSON response
func (c *Controller) CreateResponse(w http.ResponseWriter, code int, dto interface{}) {
	response, _ := json.Marshal(dto)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
