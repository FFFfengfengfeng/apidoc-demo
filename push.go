package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/push", func(context *gin.Context) {
		pusher := context.Writer.Pusher()

		if pusher != nil {
			err := pusher.Push("/assets/app.js", nil)

			if err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		context.HTML(200, "https", gin.H{
			"status": "success",
		})
	})
	r.Run(":9999")
}
