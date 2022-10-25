package models

import "time"

type SizeTotal struct {
	IDSizeTotal int       `json:"id_size_total"`
	LastValue   string    `json:"last_value"`
	TimeLoad    time.Time `json:"time_load"`
}
