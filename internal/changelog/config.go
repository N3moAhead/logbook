package changelog

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	OutputPath   string
	TemplatePath string
}

func ReadConfig() Config {
	var config Config = Config{
		OutputPath:   "./CHANGELOG.md",
		TemplatePath: "",
	}
	content, err := os.ReadFile(".logbook.json")
	if err != nil {
		return config
	}
	var payload Config
	err = json.Unmarshal(content, &payload)
	if err != nil {
		fmt.Println("Error while reading the config file: ", err)
		return config
	}
	return payload
}
