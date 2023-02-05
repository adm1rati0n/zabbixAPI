package models

import "time"

type HostWork struct {
	Hostid           int    `json:"hostid"`
	Name_host        string `json:"name_host"`
	Work_calculation string `json:"work_calculation"`
}

type HostWorkDB struct {
	Hostid                int       `json:"hostid"`
	Name_host             string    `json:"name_host"`
	Work_calculation      string    `json:"work_calculation"`
	Time_size_free        time.Time `json:"time_size_free"`
	Time_size_total       time.Time `json:"time_size_total"`
	Time_memory_total     time.Time `json:"time_memory_total"`
	Time_memory_available time.Time `json:"time_memory_available"`
	Time_cpu_util_user    time.Time `json:"time_cpu_util_user"`
}
