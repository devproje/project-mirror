package router

import (
	"fmt"
	"html/template"
	"math"
	"os"
	"strings"

	"github.com/devproje/project-mirror/src/config"
	"github.com/gin-gonic/gin"
)

type FileData struct {
	URL      string
	Name     string
	Size     string
	Modified string
}

func arrToStr(arr []FileData) *string {
	var str = ""
	for _, i := range arr {
		str += createElement(i.URL, i.Name, i.Size, i.Modified)
	}

	return &str
}

func getFileSize(size float64) string {
	var suffixes = [5]string{"Bytes", "KB", "MB", "GB", "TB"}
	if size == 0 {
		return fmt.Sprintf("%.0f %s", size, suffixes[0])
	}

	base := math.Log(float64(size)) / math.Log(1024)
	converted := math.Pow(1024, base-math.Floor(base))
	suffix := suffixes[int(math.Floor(base))]

	if size > 1023 {
		return fmt.Sprintf("%.1f %s", converted, suffix)
	}

	return fmt.Sprintf("%.0f %s", converted, suffix)
}

func createElement(path, name, size, modified string) string {
	return fmt.Sprintf(
		`<tr class="file">
			<td><a href="%s">%s</a></td>
			<td>%s</td>
			<td>%s</td>
		</tr>`,
		path,
		name,
		size,
		modified,
	)
}

func MirrorWorker(ctx *gin.Context, path string) {
	sort, asc := ctx.GetQuery("sort")
	if config.Get().Auth {
		_, status := CheckLogin(ctx)
		if status != 200 {
			ctx.Redirect(301, "/v1/login")
			return
		}
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

	if !file.IsDir() {
		ctx.FileAttachment(iPath, file.Name())
		return
	}

	dir := read(path, sort, asc)
	ctx.HTML(200, "index.html", gin.H{
		"name":     config.Get().Name,
		"dir_name": path,
		"content":  template.HTML(*dir),
	})
}

func read(path string, srt string, asc bool) *string {
	dir, _ := os.ReadDir(fmt.Sprintf(".data/%s", path))
	var back string
	var files []FileData

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

		files = append(files, FileData{
			URL:      back,
			Name:     "../",
			Size:     "-",
			Modified: "-",
		})
	}

	for _, file := range dir {
		ph := fmt.Sprintf("/%s/%s", path, file.Name())
		if path == "/" {
			ph = fmt.Sprint(file.Name())
		}

		format := "2006-01-02 03:04 PM"
		var name, size string
		finfo, _ := file.Info()
		modified := finfo.ModTime().Format(format)
		if file.IsDir() {
			name = file.Name() + "/"
			size = "-"
			modified = "-"
		} else {
			name = file.Name()
			size = getFileSize(float64(finfo.Size()))
		}

		files = append(files, FileData{
			URL:      ph,
			Name:     name,
			Size:     size,
			Modified: modified,
		})
	}

	return arrToStr(files)
}
