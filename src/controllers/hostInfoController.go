package controllers

import (
	"dbConnection"
	"github.com/gorilla/mux"
	"html/template"
	"models"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetHostInfo(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("select value_percpu_avg1, time_percpu_avg1, value_percpu_avg5, time_percpu_avg5, value_percpu_avg15,time_percpu_avg15, value_size_free, time_size_free, value_size_total, time_size_total, value_memory_available, time_memory_available, value_memory_total, time_memory_total, value_cpu_util_idle, time_value_cpu_util_idle, value_cpu_util_user, time_cpu_util_user from host_info where hostid = $1", id)
	if err != nil {
		panic(err)
	}
	var hostsInfo []models.Host_info
	for rows.Next() {
		var hostInfo models.Host_info
		err = rows.Scan(&hostInfo.Value_percpu_avg1, &hostInfo.Time_percpu_avg1, &hostInfo.Value_percpu_avg5, &hostInfo.Time_percpu_avg5,
			&hostInfo.Value_percpu_avg15, &hostInfo.Time_percpu_avg15, &hostInfo.Value_size_free, &hostInfo.Time_size_free, &hostInfo.Value_size_total,
			&hostInfo.Time_size_total, &hostInfo.Value_memory_available, &hostInfo.Time_memory_available, &hostInfo.Value_memory_total, &hostInfo.Time_memory_total,
			&hostInfo.Value_cpu_util_idle, &hostInfo.Time_value_cpu_util_idle, &hostInfo.Value_cpu_util_user, &hostInfo.Time_cpu_util_user)
		if err != nil {
			panic(err)
		}
		replacer := strings.NewReplacer("T", " ", "Z", " ")
		hostInfo.Time_memory_total = replacer.Replace(hostInfo.Time_memory_total)
		hostInfo.Time_memory_available = replacer.Replace(hostInfo.Time_memory_available)
		hostInfo.Time_cpu_util_user = replacer.Replace(hostInfo.Time_cpu_util_user)
		hostInfo.Time_value_cpu_util_idle = replacer.Replace(hostInfo.Time_value_cpu_util_idle)
		hostInfo.Time_percpu_avg1 = replacer.Replace(hostInfo.Time_percpu_avg1)
		hostInfo.Time_percpu_avg5 = replacer.Replace(hostInfo.Time_percpu_avg5)
		hostInfo.Time_percpu_avg15 = replacer.Replace(hostInfo.Time_percpu_avg15)
		hostInfo.Time_size_free = replacer.Replace(hostInfo.Time_size_free)
		hostInfo.Time_size_total = replacer.Replace(hostInfo.Time_size_total)
		hostsInfo = append(hostsInfo, hostInfo)
	}
	rows, err = db.Query("select * from host_work")
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
	var hostMap models.HostMap
	hostMap.Hosts = hosts
	hostMap.HostsInfo = hostsInfo
	dir, _ := os.Getwd()
	tmpl := template.Must(template.ParseFiles(dir + "/src/template/hostInfoTemplate.html"))
	err = tmpl.Execute(w, hostMap)
	if err != nil {
		panic(err)
	}
}
