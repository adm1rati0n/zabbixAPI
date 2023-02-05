package configHandler

import (
	"encoding/json"
	"fmt"
	"models"
	"os"
)

func GetConfig() models.ConfigModel {
	dir, _ := os.Getwd()
	file, _ := os.Open(dir + "/src/cfg.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	var config models.ConfigModel
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config)
	return config
}
