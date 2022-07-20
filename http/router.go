package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterAuthRouter(r *gin.Engine) {
	productGroup := r.Group("/api/telegraf")
	productGroup.GET("getFile", GetFile)
	productGroup.GET("getDir", GetDir)
	productGroup.PUT("update", Update)
	productGroup.DELETE("delete", Delete)
	productGroup.PUT("touch", Touch)
	productGroup.GET("exist", Exist)

}
