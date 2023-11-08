package controller

import (
	"dynamic-admission-webhook-example/service"
	"github.com/gin-gonic/gin"
)

func AlwaysDeny(context *gin.Context) {
	w := context.Writer
	r := context.Request
	denyService := service.NewAlwaysDenyService()
	serve(w, r, newDelegateToV1AdmitHandler(denyService.AlwaysDeny))
}
