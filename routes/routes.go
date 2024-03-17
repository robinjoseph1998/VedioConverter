package routes

import (
	"VedioConverter/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouters(r *gin.Engine) {
	r.POST("/mp4tomkv", controllers.Mp4ToMkvConverter)
}
