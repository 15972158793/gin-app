package main

import (
	"fmt"

	"github.com/15972158793/gin-app/controller"

	"github.com/15972158793/gin-app/dao/mysql"
	"github.com/15972158793/gin-app/dao/redis"
	"github.com/15972158793/gin-app/pkg"
	"github.com/15972158793/gin-app/router"
	"github.com/15972158793/gin-app/setting"
)

func setUp() {

	if err := setting.SetUp(); err != nil {
		fmt.Println("setting.SetUp() failed ...")
		return
	}

	pkg.SetUp()

	if err := mysql.SetUp(); err != nil {
		fmt.Println("mysql.SetUp() failed ...")
		return
	}
	defer mysql.Close()

	if err := redis.SetUp(); err != nil {
		fmt.Println("redis.SetUp() failed ...")
		return
	}
	defer redis.Close()

	// 定义错误翻译器
	controller.InitTranslator("zh")

	router.SetUp()

}

func main() {
	setUp()
}
