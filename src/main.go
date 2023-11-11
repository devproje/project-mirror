package main

import (
	"flag"
	"fmt"
	"text/template"

	"github.com/devproje/plog/level"
	"github.com/devproje/plog/log"
	"github.com/devproje/project-mirror/src/auth"
	"github.com/devproje/project-mirror/src/router"
	"github.com/gin-gonic/gin"
)

var (
	port  = 3000
	debug = false
)

func init() {
	flag.IntVar(&port, "port", 3000, "set service port")
	flag.BoolVar(&debug, "debug", false, "set debug mode")
	flag.Parse()

	log.SetLevel(level.Info)
	gin.SetMode(gin.ReleaseMode)
	if debug {
		log.SetLevel(level.Trace)
		gin.SetMode(gin.DebugMode)
	}
}

func main() {
	auth.Init()
	var url string
	app := gin.Default()
	app.SetFuncMap(template.FuncMap{})
	app.LoadHTMLGlob("pages/*.html")
	router.New(app)

	go func() {
		if !debug {
			log.Infof("service port listening for: %d\n", port)
		} else {
			log.Debugf("development service listening for: http://localhost:%d\n", port)
		}
	}()

	if !debug {
		url = fmt.Sprintf(":%d", port)
	} else {
		url = fmt.Sprintf("127.0.0.1:%d", port)
	}

	err := app.Run(url)
	if err != nil {
		log.Fatalf("current port already binding: %d\n", port)
	}
}
