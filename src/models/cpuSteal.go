package models

import "time"

type CPUSteal struct {
	IDCpuSteal        int       `json:"id_cpu_steal"`
	HostID            int       `json:"host_id"`
	ValueCpuSteal     float32   `json:"value_cpu_steal"`
	TimeValueCpuSteal time.Time `json:"time_value_cpu_steal"`
}
