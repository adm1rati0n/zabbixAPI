package main

import (
	"controllers"
	"dbConnection"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	dbConnection.DBconnect()
	r := mux.NewRouter()
	r.HandleFunc("/hosts", controllers.GetHosts).Methods("GET")
	r.HandleFunc("/hosts/{id}", controllers.GetHostInfo).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:4000", r))
}
