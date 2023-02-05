package models

type Host_info struct {
	Value_percpu_avg1        float32
	Time_percpu_avg1         string
	Value_percpu_avg5        float32
	Time_percpu_avg5         string
	Value_percpu_avg15       float32
	Time_percpu_avg15        string
	Value_size_free          string
	Time_size_free           string
	Value_size_total         string
	Time_size_total          string
	Value_memory_available   string
	Time_memory_available    string
	Value_memory_total       string
	Time_memory_total        string
	Value_cpu_util_idle      float32
	Time_value_cpu_util_idle string
	Value_cpu_util_user      float32
	Time_cpu_util_user       string
}
