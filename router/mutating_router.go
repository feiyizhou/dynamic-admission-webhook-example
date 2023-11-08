package router

import "github.com/gin-gonic/gin"

func MutatingRouterSet(group *gin.RouterGroup) {
	mutateGroup := group.Group("/mutate")
	mutateGroup.GET("")
}
