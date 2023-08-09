package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/devproje/plog/level"
	"github.com/devproje/plog/log"
	"github.com/devproje/project-mirror/src/router"
	"github.com/gin-gonic/gin"
)

var port = 3000

func init() {
	flag.IntVar(&port, "port", 3000, "set service port")
	log.SetLevel(level.Info)
	gin.SetMode(gin.ReleaseMode)
	if _, err := os.Stat(".data"); err != nil {
		err := os.Mkdir(".data", 0755)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func main() {
	app := gin.Default()
	app.SetFuncMap(template.FuncMap{})
	app.LoadHTMLGlob("static/*.html")
	router.New(app)

	go func() {
		log.Infof("service port bind: %d\n", port)
	}()
	err := app.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("current port already binding: %d\n", port)
	}
}
