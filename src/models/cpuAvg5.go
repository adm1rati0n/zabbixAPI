package models

import "time"

type CpuAvg5 struct {
	IdPerCpuAvg5 int       `json:"id_per_cpu_avg_5"`
	CpuLoad      float32   `json:"cpu_load"`
	TimeLoad     time.Time `json:"time_load"`
}
