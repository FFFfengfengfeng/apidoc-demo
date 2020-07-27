package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/reader", func(context *gin.Context) {
		res, err := http.Get("https://cn.bing.com/images/search?view=detailV2&ccid=1e3YVW94&id=1DE25DF0884299EC1E9E2F0324CE52C8CD7CACD0&thid=OIP.1e3YVW946dgy5uJH764JXwHaFj&mediaurl=http%3a%2f%2fwww.technosamrat.com%2fwp-content%2fuploads%2f2012%2f02%2fOcean-Wallpapers-Images1.jpg&exph=1200&expw=1600&q=%e5%9b%be%e7%89%87&simid=608040676184098104&ck=6B1C6B998F6A8207087EDC88D3670631&selectedIndex=1&ajaxhist=0")

		if err != nil || res.StatusCode != http.StatusOK {
			context.Status(http.StatusServiceUnavailable)
			return
		}

		reader := res.Body

		contentLength := res.ContentLength
		contentType := res.Header.Get("Content-Type")
		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}
		context.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	r.Run(":9999")
}
