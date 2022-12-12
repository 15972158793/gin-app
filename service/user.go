package service

import (
	"github.com/15972158793/gin-app/dao/mysql"
	"github.com/15972158793/gin-app/models"
	snow "github.com/15972158793/gin-app/pkg/snowflake"
)

func UserSignUp(param *models.ParamsUserSignUp) error {

	if err := mysql.CheckUserExist(param.Name); err != nil {
		return err
	}
	// 创建用户
	user := &models.User{
		UserID:   snow.GenerateID(),
		Name:     param.Name,
		Password: param.OriginalPassword,
	}
	return mysql.InsertUser(user)
}
