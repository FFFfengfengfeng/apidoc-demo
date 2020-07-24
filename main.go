package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r := gin.Default()

	r.GET("/json", func(context *gin.Context) {
		data := map[string]interface{}{
			"lang": "go语言",
			"tag": "<br>",
		}
		context.AsciiJSON(http.StatusOK, data)
	})

	r.Run(":9999")
}