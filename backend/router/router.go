package router

import (
	"backend/router/api"
	"backend/router/static"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	static.InitStatic(r)
	api.LoadAPI(r)
}
