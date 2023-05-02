package static

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitStatic(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.File("static/index.html")
	})
	r.NoRoute(static(http.Dir("static")))
}

// Custom FileSystem implementation that disables directory listing
func static(fs http.FileSystem) gin.HandlerFunc {
	fileServer := http.FileServer(fs)

	return func(c *gin.Context) {
		// Disable directory listing
		if strings.HasSuffix(c.Request.URL.Path, "/") {
			c.File("static/index.html")
			return
		}

		// Serve static file
		file := c.Request.URL.Path[1:]

		if _, err := fs.Open(file); err != nil {
			// Try to add .html to the end of the file
			file += ".html"
			c.File("static/" + file)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}
