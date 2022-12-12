package service

import (
	"github.com/15972158793/gin-app/dao/mysql"
	"github.com/15972158793/gin-app/models"
	snow "github.com/15972158793/gin-app/pkg/snowflake"
)

func UserSignUp(param *models.ParamsUserSignUp) error {

	if err := mysql.CheckUserExist(param.UserName); err != nil {
		return err
	}
	// 创建用户
	user := &models.User{
		UserID:   snow.GenerateID(),
		UserName: param.UserName,
		Password: param.OriginalPassword,
	}
	return mysql.InsertUser(user)
}

func UserLogin(param *models.ParamsUserLogin) (err error) {

	user := &models.User{
		UserID:   0,
		UserName: param.UserName,
		Password: param.Password,
	}
	return mysql.UserLogin(user)
}
