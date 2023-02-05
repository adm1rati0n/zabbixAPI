package models

import "time"

type CPUInterrupt struct {
	IDCpuInterrupt        int       `json:"id_cpu_interrupt"`
	HostID                int       `json:"host_id"`
	ValueCpuInterrupt     float32   `json:"value_cpu_interrupt"`
	TimeValueCpuInterrupt time.Time `json:"time_value_cpu_interrupt"`
}
