package api

import (
	"github.com/gin-gonic/gin"
)

func IsConnected(c *gin.Context) {
	// Check if has admin cookie
	cookie, err := c.Cookie("admin")
	if err != nil {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}

	// Check if cookie is valid
	if cookie != "Us3_m3_d4ddy" {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "connected",
	})
}
