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

func createFileElement(path string, fileinfo os.FileInfo) string {
	format := "2006-01-02 03:04 PM"
	var filename string
	modified := fileinfo.ModTime().Format(format)
	if fileinfo.IsDir() {
		filename = fileinfo.Name() + "/"
	} else {
		filename = fileinfo.Name()
	}

	return fmt.Sprintf(
		`<a class="file" href="%s">
			<p class="file_name">%s</p>
			<p class="file_size">%d bytes</p>
			<p class="file_modified">%s</p>
		</a>`,
		path,
		filename,
		fileinfo.Size(),
		modified,
	)
}

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
	var back, files string

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

		files += fmt.Sprintf(
			`<a class="file" href="%s">
				<p class="file_name">../</p>
				<p class="file_size"></p>
				<p class="file_modified"></p>
			</a>`,
			back,
		)
	}

	for _, file := range dir {
		ph := fmt.Sprintf("/%s/%s", path, file.Name())
		if path == "/" {
			ph = fmt.Sprint(file.Name())
		}

		fileinfo, _ := file.Info()
		files += createFileElement(ph, fileinfo)
	}

	return &files
}
