package router

import (
	"time"

	"github.com/devproje/plog/log"
	"github.com/devproje/project-mirror/src/auth"
	"github.com/devproje/project-mirror/src/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Login(ctx *gin.Context) {
	if _, status := CheckLogin(ctx); status == 200 {
		ctx.Redirect(301, "/")
	}

	acc := &auth.Account{
		Username: ctx.PostForm("Username"),
		Password: ctx.PostForm("Password"),
	}

	var errorForm = func() {
		ctx.HTML(401, "login.html", gin.H{
			"name":    config.Get().Name,
			"message": "Wrong username or password, Please try again",
		})
	}

	if status := acc.Login(); !status {
		errorForm()
		return
	}

	token, err := acc.GetJwtToken()
	if err != nil {
		errorForm()
		return
	}

	ctx.Header("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-cache=0, max-age=0")
	ctx.Header("Last-Modified", time.Now().String())
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "-1")

	ctx.SetCookie("access-token", token, 1800, "", "", false, false)

	ctx.Redirect(301, "/")
}

func LoginForm(ctx *gin.Context) {
	if _, status := CheckLogin(ctx); status == 200 {
		ctx.Redirect(301, "/")
	}

	ctx.HTML(404, "login.html", gin.H{
		"name": config.Get().Name,
	})
}

func CheckLogin(ctx *gin.Context) (string, int) {
	token, err := ctx.Request.Cookie("access-token")
	if err != nil {
		return "", 401
	}

	str := token.Value

	if str == "" {
		return "", 401
	}

	claims := &auth.Claims{}
	_, err = jwt.ParseWithClaims(str, claims, func(t *jwt.Token) (interface{}, error) {
		return auth.JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", 401
		}

		log.Errorln(err)
		return "", 403
	}

	return claims.UserID, 200
}
