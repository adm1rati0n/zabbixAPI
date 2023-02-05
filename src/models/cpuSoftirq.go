package models

import "time"

type CPUSofirq struct {
	IDCpuSoftirq        int       `json:"id_cpu_softirq"`
	HostID              int       `json:"host_id"`
	ValueCpuSoftirq     float32   `json:"value_cpu_softirq"`
	TimeValueCpuSoftirq time.Time `json:"time_value_cpu_softirq"`
}
