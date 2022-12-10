package mysql

import (
	"fmt"

	"github.com/15972158793/gin-app/setting"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func SetUp() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.AppConfig.MySqlConfig.User,
		setting.AppConfig.MySqlConfig.Password,
		setting.AppConfig.MySqlConfig.Host,
		setting.AppConfig.MySqlConfig.Port,
		setting.AppConfig.MySqlConfig.DB)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect(setting.AppConfig.MySqlConfig.Type, dsn)
	if err != nil {
		zap.L().Error("connect DB failed ", zap.Error(err))
		return
	}

	if err = db.Ping(); err != nil {
		zap.L().Error("DB ping failed", zap.Error(err))
		return
	}

	zap.L().Info("DB init success ...")

	db.SetMaxOpenConns(setting.AppConfig.MySqlConfig.MaxOpenConns)
	db.SetMaxIdleConns(setting.AppConfig.MySqlConfig.MaxIdleConns)
	return
}

func Close() {
	defer db.Close()
}
