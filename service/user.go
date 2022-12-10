package service

import (
	"github.com/15972158793/gin-app/dao/mysql"
	snow "github.com/15972158793/gin-app/pkg/snowflake"
)

func UserSignUp() {

	mysql.UserExistByName()

	userId := snow.GenerateID()

	mysql.InsertUser(userId)

}
