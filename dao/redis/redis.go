package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/15972158793/gin-app/setting"

	"go.uber.org/zap"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

func SetUp() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", setting.AppConfig.RedisConfig.Host, setting.AppConfig.RedisConfig.Port),
		Password:     setting.AppConfig.RedisConfig.Password, // no password set
		DB:           setting.AppConfig.RedisConfig.DB,       // use default DB
		PoolSize:     setting.AppConfig.RedisConfig.PoolSize, // 连接池大小
		MinIdleConns: setting.AppConfig.RedisConfig.MaxIdleConns,
	})

	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func Close() {
	rdb.Close()
}

func Set(key string, data interface{}, time0 time.Duration) error {
	return rdb.Set(ctx, key, data, time0).Err()
}

func Exists(key string) bool {
	_, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return false
	} else {
		return true
	}
}

func Get(key string) (val string, err error) {
	val, err = rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		zap.L().Error("key does not exist")
		return val, err
	} else if err != nil {
		zap.L().Error(fmt.Sprintf("redis get %s failed ...", key))
		return val, err
	} else {
		return val, nil
	}
	return val, nil
}
