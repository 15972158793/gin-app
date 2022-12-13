package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"

	"github.com/15972158793/gin-app/middleware"
	"github.com/15972158793/gin-app/setting"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

// SetUp 初始化
func SetUp() {

	gin.SetMode(setting.AppConfig.Mode)

	engine = gin.New()
	// 设置受信任的代理
	_ = engine.SetTrustedProxies([]string{"127.0.0.1"})
	// 设置url中的大写自动转小写，..和//自动移除，
	engine.RedirectFixedPath = true
	// 开启请求方法不允许，并且返回状态码405
	engine.HandleMethodNotAllowed = true
	// 设置允许从远程客户端的哪个header头中获取ip（需搭配设置受信任的代理一起使用）
	engine.RemoteIPHeaders = append(engine.RemoteIPHeaders, "Client-IP")

	// 使用中间件
	engine.Use(middleware.Logger(), middleware.Recovery(false))

	// 注册路由
	initRoute()

	// 启动服务
	startService()
}

// startService 平滑关闭
func startService() {
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.AppConfig.Port),
		Handler:        engine,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			zap.L().Error(fmt.Sprintf("Listen: %s\n", err))
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown:" + err.Error())
	}
	zap.L().Info("Server exiting")
}
