package router

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/devproje/plog/log"
	"github.com/devproje/project-mirror/src/config"
	"github.com/gin-gonic/gin"
)

func MirrorWorker(ctx *gin.Context, path string) {
	if config.Get().Auth {
		_, status := CheckLogin(ctx)
		if status != 200 {
			ctx.Redirect(301, "/v1/login")
			return
		}
		log.Debugln("test")
	}

	iPath := fmt.Sprintf(".data/%s", path)
	file, err := os.Stat(iPath)
	if err != nil {
		ctx.JSON(500, gin.H{
			"status": 500,
			"error":  err.Error(),
		})
		return
	}
	log.Debugln("test")

	if !file.IsDir() {
		ctx.FileAttachment(iPath, file.Name())
		return
	}
	log.Debugln("test")

	dir := ReadDir(path)
	ctx.HTML(200, "index.html", gin.H{
		"name":     config.Get().Name,
		"dir_name": path,
		"content":  template.HTML(*dir),
	})
}

func ReadDir(path string) *string {
	dir, _ := os.ReadDir(fmt.Sprintf(".data/%s", path))
	var back, items string

	if path != "/" {
		split := strings.Split(path, "/")
		split = split[:len(split)-1]

		back = "/"
		for i, p := range split {
			if i == len(split)-1 {
				back += p
				break
			}

			back += fmt.Sprintf("%s/", p)
		}

		if back == "" {
			back = "../"
		}

		items += fmt.Sprintf("<a id='item' href='%s'><p>../</p></a>\n", back)
	}

	for _, item := range dir {
		ph := fmt.Sprintf("%s/%s", path, item.Name())
		if path == "/" {
			ph = fmt.Sprintf("%s", item.Name())
		}

		if item.IsDir() {
			items += fmt.Sprintf("<a id='item' href='/%s'><p>%s/</p></a>\n", ph, item.Name())
		} else {
			items += fmt.Sprintf("<a id='item' href='/%s'><p>%s</p></a>\n", ph, item.Name())
		}
	}

	return &items
}
