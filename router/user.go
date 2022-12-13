package router

import (
	"github.com/15972158793/gin-app/controller"
	"github.com/15972158793/gin-app/middleware"
)

func registerUserRouter() {

	apiV1 := engine.Group("/api/v1/")

	apiV1.Use()
	{
		apiV1.POST("/user/sign_up", controller.UserSignUp)
		apiV1.POST("/user/login", controller.UserLogin)
	}

	// 需要token的
	apiV1.Use(middleware.JWTAuthorization())
	{
		apiV1.GET("/user/info", controller.UserInfo)
	}
}
