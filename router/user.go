package router

import (
	v1 "github.com/15972158793/gin-app/controller/v1"
	v2 "github.com/15972158793/gin-app/controller/v2"
	"github.com/15972158793/gin-app/middleware"
)

func registerUserRouter() {

	// v1版本
	apiV1 := engine.Group("/api/v1/")
	apiV1.Use()
	{
		apiV1.POST("/user/sign_up", v1.UserSignUp)
		apiV1.POST("/user/login", v1.UserLogin)
	}
	// 需要token的
	apiV1.Use(middleware.JWTAuthorization())
	{
		apiV1.GET("/user/info", v1.UserInfo)
	}

	// v2版本
	apiV2 := engine.Group("/api/v2/")
	apiV2.Use()
	{
		apiV2.POST("/user/sign_up", v2.UserSignUp)
		apiV2.POST("/user/login", v2.UserLogin)
	}
}
