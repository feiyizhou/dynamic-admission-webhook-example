package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.New()
	group := router.Group("/admission")
	group.GET("/ready", func(context *gin.Context) {
		context.JSON(200, "ok")
		return
	})
	ValidatingRouterSet(group)
	//MutatingRouterSet(group)
	return router
}
