package controllers

import (
	"dbConnection"
	"encoding/json"
	"fmt"
	"models"
	"net/http"
)

//w http.ResponseWriter, r *http.Request

func GetHosts(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	rows, err := db.Query("select * from host")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var hosts []models.Host
	for rows.Next() {
		var host models.Host
		err = rows.Scan(&host.Hostid, &host.Hostname)
		if err != nil {
			panic(err)
		}
		hosts = append(hosts, host)
	}
	fmt.Println(hosts)
	json.NewEncoder(w).Encode(hosts)
}
