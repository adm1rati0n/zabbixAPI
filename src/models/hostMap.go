package models

type HostMap struct {
	Hosts     []HostWork  `json:"hosts"`
	HostsInfo []Host_info `json:"hosts_info"`
}
