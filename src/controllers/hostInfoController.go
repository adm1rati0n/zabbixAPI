package controllers

import (
	"database/sql"
	"dbConnection"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"models"
	"net/http"
	"strconv"
)

func GetHostInfo(w http.ResponseWriter, r *http.Request) {
	var jsonResponse models.HostInfo
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	jsonResponse.CpuUtilIdle = GetCpuIdle(db, id)
	jsonResponse.CpuUtilUser = GetCpuUser(db, id)
	jsonResponse.PerCpuAvg1 = GetCpuAvg1(db, id)
	jsonResponse.PerCpuAvg5 = GetCpuAvg5(db, id)
	jsonResponse.PerCpuAvg15 = GetCpuAvg15(db, id)
	jsonResponse.SizeTotal = GetSizeTotal(db, id)
	jsonResponse.SizeFree = GetSizeFree(db, id)
	jsonResponse.MemoryAvailable = GetMemoryAvailable(db, id)
	jsonResponse.MemoryTotal = GetMemoryTotal(db, id)
	json.NewEncoder(w).Encode(jsonResponse)
}

func GetCpuAvg5(db *sql.DB, id int) []models.CpuAvg5 {
	rows, err := db.Query("select id_percpu_avg5, value_percpu_avg5, time_percpu_avg5 from percpu_avg5 where hostid = $1", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var cpuAvg5 []models.CpuAvg5
	for rows.Next() {
		var cpuAvg models.CpuAvg5
		err = rows.Scan(&cpuAvg.IdPerCpuAvg5, &cpuAvg.CpuLoad, &cpuAvg.TimeLoad)
		if err != nil {
			panic(err)
		}
		cpuAvg5 = append(cpuAvg5, cpuAvg)
	}
	fmt.Println(cpuAvg5)
	return cpuAvg5
}
func GetCpuAvg15(db *sql.DB, id int) []models.CpuAvg15 {
	rows, err := db.Query("select id_percpu_avg15, value_percpu_avg15, time_percpu_avg15 from percpu_avg15 where hostid = $1", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var cpuAvg15 []models.CpuAvg15
	for rows.Next() {
		var cpuAvg models.CpuAvg15
		err = rows.Scan(&cpuAvg.IdPerCpuAvg15, &cpuAvg.CpuLoad, &cpuAvg.TimeLoad)
		if err != nil {
			panic(err)
		}
		cpuAvg15 = append(cpuAvg15, cpuAvg)
	}
	fmt.Println(cpuAvg15)
	return cpuAvg15
}

func GetCpuAvg1(db *sql.DB, id int) []models.CpuAvg1 {
	rows, err := db.Query("select id_percpu_avg1, value_percpu_avg1, time_percpu_avg1 from percpu_avg1 where hostid = $1", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var cpuAvg1 []models.CpuAvg1
	for rows.Next() {
		var cpuAvg models.CpuAvg1
		err = rows.Scan(&cpuAvg.IdPerCpuAvg1, &cpuAvg.CpuLoad, &cpuAvg.TimeLoad)
		if err != nil {
			panic(err)
		}
		cpuAvg1 = append(cpuAvg1, cpuAvg)
	}
	fmt.Println(cpuAvg1)
	return cpuAvg1
}

func GetSizeTotal(db *sql.DB, id int) []models.SizeTotal {
	rows, err := db.Query("select id_size_total, value_size_total, time_size_total from size_total where hostid = $1", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var sizeTotal []models.SizeTotal
	for rows.Next() {
		var sizeTotalRow models.SizeTotal
		err = rows.Scan(&sizeTotalRow.IDSizeTotal, &sizeTotalRow.LastValue, &sizeTotalRow.TimeLoad)
		if err != nil {
			panic(err)
		}
		sizeTotal = append(sizeTotal, sizeTotalRow)
	}
	fmt.Println(sizeTotal)
	return sizeTotal
}

func GetSizeFree(db *sql.DB, id int) []models.SizeFree {
	rows, err := db.Query("select id_size_free, value_size_free, time_size_free from size_free where hostid = $1", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var sizeFree []models.SizeFree
	for rows.Next() {
		var sizeFreeRow models.SizeFree
		err = rows.Scan(&sizeFreeRow.IDSizeFree, &sizeFreeRow.LastValue, &sizeFreeRow.TimeLoad)
		if err != nil {
			panic(err)
		}
		sizeFree = append(sizeFree, sizeFreeRow)
	}
	fmt.Println(sizeFree)
	return sizeFree
}

func GetCpuIdle(db *sql.DB, id int) []models.CpuIdle {
	rows, err := db.Query("select id_cpu_util_idle, value_cpu_util_idle, time_value_cpu_util_idle from cpu_util_idle where hostid = $1", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var cpuIdle []models.CpuIdle
	for rows.Next() {
		var cpuIdleRow models.CpuIdle
		err = rows.Scan(&cpuIdleRow.IdCpuUtilIdle, &cpuIdleRow.LastValue, &cpuIdleRow.TimeLoad)
		if err != nil {
			panic(err)
		}
		cpuIdle = append(cpuIdle, cpuIdleRow)
	}
	fmt.Println(cpuIdle)
	return cpuIdle
}

func GetCpuUser(db *sql.DB, id int) []models.CpuUser {
	rows, err := db.Query("select id_cpu_util_user, value_cpu_util_user, time_cpu_util_user from cpu_util_user where hostid = $1", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var cpuUser []models.CpuUser
	for rows.Next() {
		var cpuUserRow models.CpuUser
		err = rows.Scan(&cpuUserRow.IdCpuUtilUser, &cpuUserRow.LastValue, &cpuUserRow.TimeLoad)
		if err != nil {
			panic(err)
		}
		cpuUser = append(cpuUser, cpuUserRow)
	}
	fmt.Println(cpuUser)
	return cpuUser
}

func GetMemoryTotal(db *sql.DB, id int) []models.MemoryTotal {
	rows, err := db.Query("select id_memory_total, value_memory_total, time_memory_total from memory_total where hostid = $1", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var memoryTotal []models.MemoryTotal
	for rows.Next() {
		var memoryTotalRow models.MemoryTotal
		err = rows.Scan(&memoryTotalRow.IdMemoryTotal, &memoryTotalRow.LastValue, &memoryTotalRow.TimeLoad)
		if err != nil {
			panic(err)
		}
		memoryTotal = append(memoryTotal, memoryTotalRow)
	}
	fmt.Println(memoryTotal)
	return memoryTotal
}

func GetMemoryAvailable(db *sql.DB, id int) []models.MemoryAvailable {
	rows, err := db.Query("select id_memory_available, value_memory_available, time_memory_available from memory_available where hostid = $1", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var memoryAvailable []models.MemoryAvailable
	for rows.Next() {
		var memoryAvailableRow models.MemoryAvailable
		err = rows.Scan(&memoryAvailableRow.IdMemoryAvailable, &memoryAvailableRow.LastValue, &memoryAvailableRow.TimeLoad)
		if err != nil {
			panic(err)
		}
		memoryAvailable = append(memoryAvailable, memoryAvailableRow)
	}
	fmt.Println(memoryAvailable)
	return memoryAvailable
}
