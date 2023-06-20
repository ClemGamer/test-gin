package router

import (
	v1 "github.com/ClemGamer/test-gin/router/api/v1"
	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.RouterGroup) {
	apiGroup := r.Group("api")
	{
		setRouters(apiGroup)
	}
}

func setRouters(r *gin.RouterGroup) {

	v1.UserRouter(r)

}
