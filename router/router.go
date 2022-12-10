package router

import (
	"fmt"

	"github.com/15972158793/gin-app/controller"
	"github.com/15972158793/gin-app/middleware"
	"github.com/15972158793/gin-app/setting"
	"github.com/gin-gonic/gin"

	_ "github.com/15972158793/gin-app/docs" //注意导入
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUp() {

	gin.SetMode(setting.AppConfig.Mode)

	r := gin.New()
	r.Use(middleware.Logger(), middleware.Recovery(true))

	// 编写自动生成API文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 注册路由
	apiV1 := r.Group("/api/v1/")
	// 使用中间件
	apiV1.Use()
	{
		apiV1.POST("/user/sign_up", controller.UserSignUp)

		//apiV1.GET("/tags", v1.GetTags)
		//apiV1.POST("/tags", v1.AddTag)
		//apiV1.PUT("/tags/:id", v1.EditTag)
		//apiV1.DELETE("/tags/:id", v1.DeleteTag)
		//
		//apiV1.GET("/articles", v1.GetArticleList)
		//apiV1.GET("/article/:id", v1.GetArticle)
		//apiV1.POST("/article", v1.AddArticle)
		//apiV1.PUT("/article/:id", v1.EditArticle)
		//apiV1.DELETE("/article/:id", v1.DeleteArticle)
	}

	r.Run(fmt.Sprintf(":%d", setting.AppConfig.Port))

}
