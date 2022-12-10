package pkg

import (
	"fmt"

	"github.com/15972158793/gin-app/pkg/logger"
	snow "github.com/15972158793/gin-app/pkg/snowflake"

	"go.uber.org/zap"
)

func SetUp() {

	if err := logger.SetUp(); err != nil {
		fmt.Println("logger.SetUp() failed ...")
		return
	}

	if err := snow.SetUp("2023-01-01", 1); err != nil {
		zap.L().Error("snow.SetUp failed ...")
	}

}
