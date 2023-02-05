package models

import "time"

type HostGraph struct {
	Value_percpu_avg1        float32
	Time_percpu_avg1         time.Time
	Value_percpu_avg5        float32
	Time_percpu_avg5         time.Time
	Value_percpu_avg15       float32
	Time_percpu_avg15        time.Time
	Value_size_free          string
	Time_size_free           time.Time
	Value_size_total         string
	Time_size_total          time.Time
	Value_memory_available   string
	Time_memory_available    time.Time
	Value_memory_total       string
	Time_memory_total        time.Time
	Value_cpu_util_idle      float32
	Time_value_cpu_util_idle time.Time
	Value_cpu_util_user      float32
	Time_cpu_util_user       time.Time
}
