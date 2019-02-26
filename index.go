package main

import (
	"net/http"

	api "./api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/Users", api.GetAllUsers).Methods("GET")
	r.HandleFunc("/api/Users/{id}", api.GetUser).Methods("GET")
	r.HandleFunc("/api/Users", api.PostUser).Methods("POST")
	r.HandleFunc("/api/Users/{id}", api.DeleteUser).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
