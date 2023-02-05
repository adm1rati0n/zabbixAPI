package models

import "time"

type CpuUser struct {
	IdCpuUtilUser int       `json:"id_cpu_util_user"`
	LastValue     string    `json:"last_value"`
	TimeLoad      time.Time `json:"time_load"`
}
