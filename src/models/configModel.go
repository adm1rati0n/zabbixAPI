package models

type ConfigModel struct {
	DbHost         string `json:"host"`
	Port           int    `json:"port"`
	User           string `json:"user"`
	Password       string `json:"password"`
	DbName         string `json:"db_name"`
	RefreshSeconds string `json:"refresh_seconds"`
	ServerUrl      string `json:"server_url"`
}
