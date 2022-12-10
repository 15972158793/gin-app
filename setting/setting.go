package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type App struct {
	Name    string `mapstructure:"Name"`
	Mode    string `mapstructure:"Mode"`
	Port    int    `mapstructure:"Port"`
	Version string `mapstructure:"Version"`

	ImagePrefixUrl string `mapstructure:"ImagePrefixUrl"`
	ImageSavePath  string `mapstructure:"ImageSavePath"`
	ImageMaxSize   string `mapstructure:"ImageMaxSize"`
	ImageAllowExt  string `mapstructure:"ImageAllowExt"`

	*LogConfig   `mapstructure:"log"`
	*MySqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

var AppConfig = new(App)

type LogConfig struct {
	Level      string `mapstructure:"Level"`
	Filename   string `mapstructure:"Filename"`
	MaxSize    int    `mapstructure:"MaxSize"`
	MaxAge     int    `mapstructure:"MaxAge"`
	MaxBackups int    `mapstructure:"MaxBackups"`
}

type MySqlConfig struct {
	Type         string `mapstructure:"Type"`
	User         string `mapstructure:"User"`
	Password     string `mapstructure:"Password"`
	Host         string `mapstructure:"Host"`
	Port         int    `mapstructure:"Port"`
	DB           string `mapstructure:"DB"`
	TablePrefix  string `mapstructure:"TablePrefix"`
	MaxOpenConns int    `mapstructure:"MaxOpenConns"`
	MaxIdleConns int    `mapstructure:"MaxIdleConns"`
}

type RedisConfig struct {
	Host         string `mapstructure:"Host"`
	Port         int    `mapstructure:"Port"`
	Password     string `mapstructure:"Password"`
	DB           int    `mapstructure:"DB"`
	PoolSize     int    `mapstructure:"PoolSize"`
	MaxOpenConns int    `mapstructure:"MaxOpenConns"`
	MaxIdleConns int    `mapstructure:"MaxIdleConns"`
	Timeout      int    `mapstructure:"Timeout"`
}

func SetUp() (err error) {

	viper.SetConfigName("config")
	// 设置远程配置读取(etcd等配置中心需要)
	viper.SetConfigType("yaml")

	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("setting.go: viper.ReadInConfig failed ...,%v\n", err.Error())
		return err
	}

	if err := viper.Unmarshal(AppConfig); err != nil {
		fmt.Println("viper.Unmarshal failed ...")
	}

	// fmt.Printf("%#v,%#v,%#v", AppConfig.LogConfig, AppConfig.MySqlConfig, AppConfig.RedisConfig)

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("文件配置已经发生了修改......")
		if err := viper.Unmarshal(AppConfig); err != nil {
			fmt.Println("viper.Unmarshal update failed ...")
		}
	})
	return
}
