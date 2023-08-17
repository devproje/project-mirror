package config

import (
	"encoding/json"
	"os"

	"github.com/devproje/plog/log"
)

type Config struct {
	Name string `json:"name"`
}

func Get() *Config {
	file, _ := os.Open("server.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	data := Config{}

	err := decoder.Decode(&data)
	if err != nil {
		log.Fatalln(err)
	}

	return &data
}
