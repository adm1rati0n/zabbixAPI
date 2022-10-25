package models

type HostInfo struct {
	CpuUtilIdle     []CpuIdle         `json:"cpu_util_idle"`
	CpuUtilUser     []CpuUser         `json:"cpu_util_user"`
	PerCpuAvg1      []CpuAvg1         `json:"per_cpu_avg_1"`
	PerCpuAvg5      []CpuAvg5         `json:"per_cpu_avg_5"`
	PerCpuAvg15     []CpuAvg15        `json:"per_cpu_avg_15"`
	SizeFree        []SizeFree        `json:"size_free"`
	SizeTotal       []SizeTotal       `json:"size_total"`
	MemoryAvailable []MemoryAvailable `json:"memory_available"`
	MemoryTotal     []MemoryTotal     `json:"memory_total"`
}
