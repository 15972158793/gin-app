package mysql

import (
	"database/sql"
	"errors"

	"github.com/15972158793/gin-app/models"
	"github.com/15972158793/gin-app/pkg/utils"
)

func CheckUserExist(name string) (err error) {
	sqlStr := `select count(user_id) from user where user_name = ?`
	var count = 0
	if err = db.Get(&count, sqlStr, name); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

func UserLogin(user *models.User) (err error) {

	oldPassword := utils.MD5(user.Password)

	sqlStr := `select user_id,user_name,password from user where user_name = ?`

	err = db.Get(user, sqlStr, user.UserName)
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	if err != nil {
		return err
	}
	if oldPassword != user.Password {
		return errors.New("密码错误")
	}
	return

}

// InsertUser 插入数据
func InsertUser(user *models.User) error {
	// 加密
	password := utils.MD5(user.Password)
	sqlStr := `insert into user(user_id,user_name,password) values(?,?,?)`
	_, err := db.Exec(sqlStr, user.UserID, user.UserName, password)
	if err != nil {
		return err
	}
	return nil
}
