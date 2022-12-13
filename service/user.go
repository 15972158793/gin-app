package service

import (
	"github.com/15972158793/gin-app/dao/mysql"
	"github.com/15972158793/gin-app/models"
	"github.com/15972158793/gin-app/pkg/jwt"
	snow "github.com/15972158793/gin-app/pkg/snowflake"
	"github.com/gin-gonic/gin"
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

func UserLogin(param *models.ParamsUserLogin) (interface{}, error) {
	user := &models.User{
		UserID:   0,
		UserName: param.UserName,
		Password: param.Password,
	}
	if err := mysql.UserLogin(user); err != nil {
		return nil, err
	}
	// 登录成功后更新了user
	// 生成token
	token, err := jwt.GenerateToken(user.UserID, user.UserName)

	result := gin.H{
		"user_id":   user.UserID,
		"user_name": user.UserName,
		"token":     token,
	}
	return result, err
}
