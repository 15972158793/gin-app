package router

import (
	"github.com/15972158793/gin-app/docs"
	"github.com/15972158793/gin-app/setting"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/15972158793/gin-app/docs" //注意导入
)

// registerSwaggerRouter 注册swagger文档
func registerSwaggerRouter() {

	docs.SwaggerInfo.Host = setting.AppConfig.Host
	//docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Version = setting.AppConfig.Version

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
