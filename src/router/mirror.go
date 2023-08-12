package router

import (
	"fmt"

	"github.com/devproje/project-mirror/src/util"
	"github.com/gin-gonic/gin"
)

func mirror(ctx *gin.Context) {
	util.MirrorWorker(ctx, "/")
}

func mirrorPath(ctx *gin.Context) {
	origin := ctx.Param("path")
	child := ctx.Param("child")

	var path = origin
	if child != "/" {
		path = fmt.Sprintf("%s%s", origin, child)
	}

	util.MirrorWorker(ctx, path)
}
