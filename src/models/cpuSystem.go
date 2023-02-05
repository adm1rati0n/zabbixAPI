package models

import "time"

type CPUSystem struct {
	IDCpuSystem        int       `json:"id_cpu_system"`
	HostID             int       `json:"host_id"`
	ValueCpuSystem     float32   `json:"value_cpu_system"`
	TimeValueCpuSystem time.Time `json:"time_value_cpu_system"`
}
