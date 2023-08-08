package main

import (
	"os"

	"github.com/devproje/plog/level"
	"github.com/devproje/plog/log"
	"github.com/gin-gonic/gin"
)

func init() {
	log.SetLevel(level.Info)
	if _, err := os.Stat(".data"); err != nil {
		err := os.Mkdir(".data", 0755)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func main() {
	app := gin.Default()
	app.Static("/file", ".data")

	app.Run(":3000")
}
