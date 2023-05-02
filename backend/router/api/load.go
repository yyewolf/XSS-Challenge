package api

import (
	"backend/router/middlewares"

	"github.com/gin-gonic/gin"
)

func LoadAPI(r *gin.Engine) {
	gr := r.Group("/api")
	gr.GET("/connected", IsConnected)
	gr.POST("/report", NewReport)

	// Admin routes
	gr.GET("/reports", middlewares.IsConnected, ListReports)
	gr.DELETE("/report/:id", middlewares.IsConnected, DelReport)
	gr.GET("/report/:id", middlewares.IsConnected, GetReport)
	gr.GET("/flag", middlewares.IsConnected, GetFlag)
}
