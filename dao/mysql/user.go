package mysql

import (
	"errors"

	"github.com/15972158793/gin-app/models"
	"github.com/15972158793/gin-app/pkg/utils"
)

func CheckUserExist(name string) (err error) {
	sqlStr := `select count(user_id) from user where name = ?`
	var count = 0
	if err = db.Get(&count, sqlStr, name); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

// InsertUser 插入数据
func InsertUser(user *models.User) error {
	// 加密
	password := utils.MD5(user.Password)
	sqlStr := `insert into user(user_id,name,password) values(?,?,?)`
	_, err := db.Exec(sqlStr, user.UserID, user.Name, password)
	if err != nil {
		return err
	}
	return nil
}
