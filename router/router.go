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
	// 设置受信任的代理
	r.SetTrustedProxies([]string{"127.0.0.1"})
	// 设置url中的大写自动转小写，..和//自动移除，
	r.RedirectFixedPath = true
	// 开启请求方法不允许，并且返回状态码405
	r.HandleMethodNotAllowed = true
	// 设置允许从远程客户端的哪个header头中获取ip（需搭配设置受信任的代理一起使用）
	r.RemoteIPHeaders = append(r.RemoteIPHeaders, "Client-IP")

	r.Use(middleware.Logger(), middleware.Recovery(false))

	// 编写自动生成API文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 注册路由
	apiV1 := r.Group("/api/v1/")
	// 使用中间件
	apiV1.Use()
	{
		apiV1.POST("/user/sign_up", controller.UserSignUp)
		apiV1.POST("/user/login", controller.UserLogin)

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
