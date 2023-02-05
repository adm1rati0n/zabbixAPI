package models

import "time"

type CpuAvg1 struct {
	IdPerCpuAvg1 int       `json:"id_per_cpu_avg_1"`
	CpuLoad      float32   `json:"cpu_load"`
	TimeLoad     time.Time `json:"time_load"`
}
