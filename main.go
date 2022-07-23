package main

import (
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"fmt"
	_ "github.com/eojan-tec/telegraf-sidecar/docs"
	"github.com/eojan-tec/telegraf-sidecar/http"
	"github.com/eojan-tec/telegraf-sidecar/model"
	"github.com/gin-gonic/gin"
	"github.com/nooocode/pkg/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title Telegraf API
// @version 0.0.1
// @description xinagri telegraf server
// @description 错误代码
// @description Success=20000,InternalServerError=50000,BadRequest=40000,Unauthorized=40001,ErrRecordNotFound=40002,UserNameOrPasswordIsWrong=41001,UserIsNotExist=41002,NoPermission=41003,TokenInvalid=41004,TokenExpired=41005,UserDisabled=41006
//
//@securityDefinitions.basic BasicAuth
//@securityDefinitions.apikey ApiKeyAuth
//@in header
//@name Authorization
//
//@securitydefinitions.oauth2.application OAuth2Application
//@tokenUrl https://example.com/oauth/token
//@scope.write Grants write access
//@scope.admin Grants read and write access to administrative information
//
//@securitydefinitions.oauth2.implicit OAuth2Implicit
//@authorizationUrl https://example.com/oauth/authorize
//@scope.write Grants write access
//@scope.admin Grants read and write access to administrative information
//
//@securitydefinitions.oauth2.password OAuth2Password
//@tokenUrl https://example.com/oauth/token
//@scope.read Grants read access
//@scope.write Grants write access
//@scope.admin Grants read and write access to administrative information
//
//@securitydefinitions.oauth2.accessCode OAuth2AccessCode
//@tokenUrl https://example.com/oauth/token
//@authorizationUrl https://example.com/oauth/authorize
//@scope.admin Grants read and write access to administrative information
func main() {
	model.Init(nil)
	log.Println("started server")
	Start(48087)
}
func Start(port int) {
	r := gin.Default()
	r.Use(utils.Cors())
	//开启权限校验
	// r.Use(middleware.AuthRequiredWithRPC)
	http.RegisterAuthRouter(r)
	r.GET("/swagger/telegraf/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(fmt.Sprintf(":%d", port))
}
