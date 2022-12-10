package main

import (
	"fmt"

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

	if err := redis.SetUp(); err != nil {
		fmt.Println("redis.SetUp() failed ...")
		return
	}

	router.SetUp()
}

func main() {
	setUp()
}
