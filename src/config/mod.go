package config

import (
	"encoding/json"
	"os"

	"github.com/devproje/plog/log"
)

const (
	filename   = "server.json"
	serverConf = `{
	"name": "Project_IO's Mirror"
}`
)

type Config struct {
	Name string `json:"name"`
}

func init() {
	if _, err := os.Stat(filename); err != nil {
		log.Errorf("`%s` is not founded, create new one...\n", filename)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatalln(err)
		}

		_, err = file.Write([]byte(serverConf))
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func Get() *Config {
	file, _ := os.Open(filename)
	defer file.Close()

	decoder := json.NewDecoder(file)
	data := Config{}

	err := decoder.Decode(&data)
	if err != nil {
		log.Fatalln(err)
	}

	return &data
}
