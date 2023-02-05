package models

import "time"

type CpuIdle struct {
	IdCpuUtilIdle int       `json:"id_cpu_util_idle"`
	LastValue     string    `json:"last_value"`
	TimeLoad      time.Time `json:"time_load"`
}
