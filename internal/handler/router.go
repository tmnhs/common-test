package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tmnhs/common-test/internal/middlerware"
)

func RegisterRouters(r *gin.Engine) {

	configRoute(r)

	configNoRoute(r)
}

func configRoute(r *gin.Engine) {

	hello := r.Group("/ping")
	{
		hello.GET("", func(c *gin.Context) {
			c.JSON(200, "pong")
		})
	}

	base := r.Group("")
	{
		base.POST("register", defaultUserRouter.Register)
		base.POST("login", defaultUserRouter.Login)
	}

	user := r.Group("/user")
	user.Use(middlerware.JWTAuth())
	{
		user.POST("del", defaultUserRouter.Delete)
		user.POST("update", defaultUserRouter.Update)
		user.POST("change_pw", defaultUserRouter.ChangePassword)
		user.GET("find", defaultUserRouter.FindById)
		user.POST("search", defaultUserRouter.Search)
	}
	upload := r.Group("/file")
	//upload.Use(middlerware.JWTAuth())
	{
		upload.POST("del", defaultUploadRouter.DeleteFile)
		upload.POST("upload", defaultUploadRouter.UploadFile)
	}
}

func configNoRoute(r *gin.Engine) {
	/*	r.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
		r.StaticFile("favicon.ico", "./dist/favicon.ico")
		r.Static("/css", "./dist/css")
		r.Static("/fonts", "./dist/fonts")
		r.Static("/js", "./dist/js")
		r.Static("/img", "./dist/img")
		r.StaticFile("/", "./dist/index.html") // 前端网页入口页面*/
}
