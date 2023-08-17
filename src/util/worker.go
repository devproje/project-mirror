package util

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/devproje/project-mirror/src/config"
	"github.com/gin-gonic/gin"
)

func MirrorWorker(ctx *gin.Context, path string) {
	var targetPath = path
	c, err := dirList(fmt.Sprintf(".data/%s", targetPath))
	if err != nil {
		errorHandler(ctx, err)
	}

	if targetPath == "/" {
		c, err = dirList(".data")
		if err != nil {
			errorHandler(ctx, err)
			return
		}
	}

	if targetPath[0] != '/' {
		targetPath = "/" + targetPath
	}

	ctx.HTML(200, "index.html", gin.H{
		"name":     config.Get().Name,
		"dir_name": targetPath,
		"content":  template.HTML(*c),
	})
}

func errorHandler(ctx *gin.Context, err error) {
	var status = 500
	ctx.JSON(status, gin.H{
		"status": status,
		"reason": err.Error(),
	})
}

func dirList(path string) (*string, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var items = ""
	if path != ".data" {
		var back = strings.ReplaceAll(path, ".data", "")
		split := strings.Split(back, "/")
		split = split[:len(split)-1]

		back = ""
		for i, j := range split {
			if i == len(split)-1 {
				back += j
				break
			}

			back += fmt.Sprintf("%s/", j)
		}

		if back == "" {
			back = "../"
		}

		items += fmt.Sprintf("<a id='item' href='%s'><p>../</p></a>\n", back)
	}

	for _, i := range dir {
		ph := strings.ReplaceAll(fmt.Sprintf("%s/%s", path, i.Name()), ".data/", "")
		if i.IsDir() {
			items += fmt.Sprintf("<a id='item' href='/%s'><p>%s/</p></a>\n", ph, i.Name())
		} else {
			items += fmt.Sprintf("<a id='item' href='/file/%s'><p>%s</p></a>\n", ph, i.Name())
		}
	}

	return &items, nil
}
