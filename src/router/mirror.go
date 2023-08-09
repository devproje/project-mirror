package router

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func mirror(ctx *gin.Context) {
	mirrorWorker(ctx, "/")
}

func mirrorPath(ctx *gin.Context) {
	path, _ := ctx.Params.Get("path")
	mirrorWorker(ctx, path)
}

func mirrorWorker(ctx *gin.Context, path string) {
	var status = 200
	var targetPath = path

	c, err := dirList(fmt.Sprintf(".data/%s", targetPath))
	if err != nil {
		status = 500
		ctx.JSON(status, gin.H{
			"status": status,
			"reason": err.Error(),
		})

		return
	}
	if targetPath == "/" {
		c, err = dirList(".data")
		if err != nil {
			status = 500
			ctx.JSON(status, gin.H{
				"status": status,
				"reason": err.Error(),
			})

			return
		}
	}

	if targetPath[0] != '/' {
		targetPath = "/" + targetPath
	}

	ctx.HTML(status, "index.html", gin.H{
		"dir_name": targetPath,
		"content":  template.HTML(*c),
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
		items += "<a id='item' href='../'><p>../</p></a>\n"
	}

	for _, i := range dir {
		ph := strings.ReplaceAll(fmt.Sprintf("%s/%s", path, i.Name()), ".data/", "")
		if i.IsDir() {
			items += fmt.Sprintf("<a id='item' href='/%s'><p>%s/</p></a>\n", ph, i.Name())
			break
		}

		items += fmt.Sprintf("<a id='item' href='/file/%s'><p>%s</p></a>\n", ph, i.Name())
	}

	return &items, nil
}
