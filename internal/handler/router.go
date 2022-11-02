package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tmnhs/common-test/internal/middlerware"
	"github.com/tmnhs/common-test/internal/model/resp"
)

func RegisterRouters(r *gin.Engine) {
	r.Use(middlerware.Cors())

	configRoute(r)

	configNoRoute(r)
}

func configRoute(r *gin.Engine) {

	hello := r.Group("/ping")
	{
		hello.GET("", func(c *gin.Context) {
			c.JSON(200, "pong")
		})
		hello.POST("", func(c *gin.Context) {
			type Hello struct {
				Name string `json:"name" form:"name"`
			}
			var h Hello
			var err error
			err = c.ShouldBindJSON(&h)
			if err != nil {
				c.JSON(resp.ERROR, err.Error())
			}
			c.JSON(200, "hello,"+h.Name)
		})
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
