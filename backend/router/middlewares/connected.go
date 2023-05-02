package middlewares

import (
	"github.com/gin-gonic/gin"
)

func IsConnected(c *gin.Context) {
	// Check if has admin cookie
	cookie, err := c.Cookie("admin")
	if err != nil {
		c.Redirect(302, "/")
		c.Abort()
		return
	}

	// Check if cookie is valid
	if cookie != "ckjsdmldMLKJSQ456" {
		c.Redirect(302, "/")
		c.Abort()
		return
	}

	c.Next()
}
