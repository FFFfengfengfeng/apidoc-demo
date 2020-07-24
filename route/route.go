package route

import (
	"github.com/gin-gonic/gin"
	"../db"
	"../controller/cate"
)

func Register(engine *gin.Engine) {
	conn := db.Connect()
	engine.GET("/cate", func(context *gin.Context) {
		res := cate.GetList(conn)
		context.JSON(200, gin.H{
			"data": res,
		})
	})
}
