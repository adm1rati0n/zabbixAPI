package models

import "time"

type MemoryAvailable struct {
	IdMemoryAvailable int       `json:"id_memory_available"`
	LastValue         string    `json:"last_value"`
	TimeLoad          time.Time `json:"time_load"`
}
