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
