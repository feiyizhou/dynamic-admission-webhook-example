package router

import (
	"dynamic-admission-webhook-example/controller"
	"github.com/gin-gonic/gin"
)

func ValidatingRouterSet(group *gin.RouterGroup) {
	validateGroup := group.Group("/validate")
	validateGroup.GET("/always-deny", controller.AlwaysDeny)
}
