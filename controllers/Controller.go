package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Controller : Base class for controllers
type Controller struct {
	Db *sql.DB
}

// ParseRequestToMap : Generic way for all controllers to parse JSON request to map
func (c *Controller) ParseRequestToMap(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	var dto map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	if decoder.Decode(&dto) != nil {
		c.CreateResponse(w, http.StatusBadRequest, "Invalid request payload")
		return nil
	}
	defer r.Body.Close()

	return dto
}

// ParseRequestToString : Generic way for all controllers to parse JSON request to string
func (c *Controller) ParseRequestToString(w http.ResponseWriter, r *http.Request) *string {

	dto, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.CreateResponse(w, http.StatusBadRequest, "Invalid request payload")
		return nil
	}
	defer r.Body.Close()

	dtoString := string(dto)

	return &dtoString
}

// CreateResponse : Generic way for all controllers to create JSON response
func (c *Controller) CreateResponse(w http.ResponseWriter, code int, dto interface{}) {
	response, _ := json.Marshal(dto)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
