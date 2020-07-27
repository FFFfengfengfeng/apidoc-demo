package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name": "index",
		})
	})
	r.GET("/cate", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name": "cate",
		})
	})
	r.GET("/*", func(context *gin.Context) {
		context.Request.URL.Path = "/"
		r.HandleContext(context)
	})
	r.Run(":9999")
}
