package config

import (
	"os"
	"os/exec"

	"github.com/devproje/plog/log"
	"github.com/devproje/project-mirror/src/util"
)

const (
	filename   = "server.json"
	serverConf = `{
	"name": "Project_IO's Mirror",
	"auth": false
}`
)

type Config struct {
	Name      string `json:"name"`
	Auth      bool   `json:"auth"`
	SecretKey string
}

func init() {
	if _, err := os.Stat(filename); err != nil {
		log.Errorf("`%s` is not founded, create new one...\n", filename)

		err = os.WriteFile(filename, []byte(serverConf), 0655)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if err := util.CreateDir(".data"); err != nil {
		log.Fatalln(err)
	}

	if err := util.CreateDir(".tmp"); err != nil {
		log.Fatalln(err)
	}
}

func Get() *Config {
	data, err := util.ParseJSON[Config](filename)
	if err != nil {
		log.Fatalln(err)
	}

	key, err := secretKey()
	if err != nil {
		log.Fatalln(err)
	}

	data.SecretKey = key
	return data
}

func secretKey() (string, error) {
	key, err := os.ReadFile(".tmp/secret.txt")
	if err != nil {
		cmd := exec.Command("openssl", "rand", "-hex", "64")
		gen, err := cmd.Output()
		if err != nil {
			return "", err
		}
		defer cmd.Process.Kill()

		err = os.WriteFile(".tmp/secret.txt", gen, 0655)
		if err != nil {
			return "", err
		}

		return string(gen), nil
	}

	return string(key), nil
}
