package router

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/15972158793/gin-app/docs" //注意导入
)

// registerSwaggerRouter 注册swagger文档
func registerSwaggerRouter() {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
