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

// ParseRequest : Generic way for all controllers to parse JSON request
func (c *Controller) ParseRequest(w http.ResponseWriter, r *http.Request, dto *interface{}) {
	decoder := json.NewDecoder(r.Body)
	if decoder.Decode(dto) != nil {
		c.CreateResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
}

// CreateResponse : Generic way for all controllers to create JSON response
func (c *Controller) CreateResponse(w http.ResponseWriter, code int, dto interface{}) {
	response, _ := json.Marshal(dto)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
