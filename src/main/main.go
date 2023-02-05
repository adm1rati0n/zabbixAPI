package main

import (
	"controllers"
	"database/sql"
	"dbConnection"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var database *sql.DB

func main() {
	dbConnection.DBconnect()
	r := mux.NewRouter()
	dir, _ := os.Getwd()
	fmt.Println(dir)
	fileServer := http.FileServer(http.Dir(dir + "/src/images/"))
	r.Handle("/src/images/", http.StripPrefix("/src/images/", fileServer))
	r.HandleFunc("/hosts", controllers.GetHosts)
	r.HandleFunc("/hosts/{id}", controllers.GetHostInfo)
	log.Fatal(http.ListenAndServe(dbConnection.Config.ServerUrl, r))
}
