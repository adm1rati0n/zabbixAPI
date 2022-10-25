package models

import "time"

type MemoryTotal struct {
	IdMemoryTotal int       `json:"id_memory_total"`
	LastValue     string    `json:"last_value"`
	TimeLoad      time.Time `json:"time_load"`
}
