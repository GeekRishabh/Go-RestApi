package api

import (
	"net/http"

	"encoding/json"
	"fmt"

	db "../db"
	"github.com/gorilla/mux"
	// "github.com/mycodesmells/mongo-go-api/db"
)

func handleError(err error, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}

// GetAllUsers returns a list of all database Users to the response.
func GetAllUsers(w http.ResponseWriter, req *http.Request) {
	rs, err := db.GetAll()
	if err != nil {
		handleError(err, "Failed to load database Users: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to load marshal data: %v", w)
		return
	}

	w.Write(bs)
}

// GetUser returns a single database User matching given ID parameter.
func GetUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	rs, err := db.GetOne(id)
	if err != nil {
		handleError(err, "Failed to read database: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to marshal data: %v", w)
		return
	}

	w.Write(bs)
}

// PostUser saves an User (form data) into the database.
func PostUser(w http.ResponseWriter, req *http.Request) {
	ID := req.FormValue("id")
	nameStr := req.FormValue("name")
	name := string(nameStr)

	user := db.User{ID: ID, Name: name}

	db.Save(user)

	w.Write([]byte("OK"))
}

// DeleteUser removes a single User (identified by parameter) from the database.
func DeleteUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	if err := db.Remove(id); err != nil {
		handleError(err, "Failed to remove User: %v", w)
		return
	}

	w.Write([]byte("OK"))
}
