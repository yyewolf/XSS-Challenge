package api

import "github.com/gin-gonic/gin"

func GetFlag(c *gin.Context) {
	c.JSON(200, gin.H{
		"flag": "HTN{XSS_1S_4W3S0M3}",
	})
}
