package router

import (
	"fmt"
	"net/http"

	"github.com/15972158793/gin-app/middleware"
	"github.com/15972158793/gin-app/setting"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUp() {

	gin.SetMode(setting.AppConfig.Mode)

	r := gin.New()
	r.Use(middleware.Logger(), middleware.Recovery(true))

	// 编写自动生成API文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mag": "啊啊啊啊",
		})
	})
	r.Run(fmt.Sprintf(":%d", setting.AppConfig.Port))

}
