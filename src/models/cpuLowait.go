package models

import "time"

type CPULowait struct {
	IDCpuLowait        int       `json:"id_cpu_lowait"`
	HostID             int       `json:"host_id"`
	ValueCpuLowait     float32   `json:"value_cpu_lowait"`
	TimeValueCpuLowait time.Time `json:"time_value_cpu_lowait"`
}
