package models

import "time"

type CpuAvg15 struct {
	IdPerCpuAvg15 int       `json:"id_per_cpu_avg_15"`
	CpuLoad       float32   `json:"cpu_load"`
	TimeLoad      time.Time `json:"time_load"`
}
