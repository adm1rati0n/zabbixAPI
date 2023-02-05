package controllers

import (
	"dbConnection"
	"html/template"
	"models"
	"net/http"
	"os"
)

//w http.ResponseWriter, r *http.Request

func GetHosts(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	rows, err := db.Query("select * from host_work")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var hosts []models.HostWork
	for rows.Next() {
		var hostDB models.HostWorkDB
		err = rows.Scan(&hostDB.Hostid, &hostDB.Name_host, &hostDB.Work_calculation,
			&hostDB.Time_size_free, &hostDB.Time_size_total, &hostDB.Time_memory_total,
			&hostDB.Time_memory_available, &hostDB.Time_cpu_util_user)
		if err != nil {
			panic(err)
		}
		var host models.HostWork
		host.Hostid = hostDB.Hostid
		host.Name_host = hostDB.Name_host
		host.Work_calculation = hostDB.Work_calculation
		hosts = append(hosts, host)
	}
	dir, _ := os.Getwd()
	tmpl := template.Must(template.ParseFiles(dir + "/src/template/hostTemplate.html"))
	err = tmpl.Execute(w, hosts)
	if err != nil {
		panic(err)
	}
}
