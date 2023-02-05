package models

type HostMapGraph struct {
	Hosts     []HostWork  `json:"hosts"`
	HostsInfo []HostGraph `json:"hosts_info"`
}
