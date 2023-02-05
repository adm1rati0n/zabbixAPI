package models

import "time"

type CPUNice struct {
	IDCpuNice        int       `json:"id_cpu_nice"`
	HostID           int       `json:"host_id"`
	ValueCpuNice     float32   `json:"value_cpu_nice"`
	TimeValueCpuNice time.Time `json:"time_value_cpu_nice"`
}
