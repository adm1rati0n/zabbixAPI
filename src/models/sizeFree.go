package models

import "time"

type SizeFree struct {
	IDSizeFree int       `json:"id_size_free"`
	LastValue  string    `json:"last_value"`
	TimeLoad   time.Time `json:"time_load"`
}
