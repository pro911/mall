package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局变量 用来保存程序的所有配置信息
var Conf = new(Config)

type Config struct {
	*AppConfig  `mapstructure:"app"`
	*HttpServer `mapstructure:"http_server"`
	*LogConfig  `mapstructure:"log"`
	//*MySQLConfig `mapstructure:"mysql"`
	*SQLiteConfig `mapstructure:"sqlite"`
	*RedisConfig  `mapstructure:"redis"`
	//*MongoConfig `mapstructure:"mongo"`
}

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	JwtSecret string `mapstructure:"jwt_secret"`
	Resources string `mapstructure:"resources"`
}

type HttpServer struct {
	Port int `mapstructure:"port"`
}

type LogConfig struct {
	Level       string `mapstructure:"level"`
	Filename    string `mapstructure:"filename"`
	ErrFilename string `mapstructure:"err_filename"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxAge      int    `mapstructure:"max_age"`
	MaxBackups  int    `mapstructure:"max_backups"`
	CloseStdout bool   `mapstructure:"close_stdout"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"db_name"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type SQLiteConfig struct {
	Host         string `mapstructure:"host"`
	DBName       string `mapstructure:"db_name"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DBName   int    `mapstructure:"db_name"`
	PoolSize int    `mapstructure:"pool_size"`
}

type MongoConfig struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	DBName     string `mapstructure:"db_name"`
	AuthSource string `mapstructure:"auth_source"`
}

func Init(configFile string) (err error) {
	if len(configFile) > 0 {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
	}

	//读取配置信息
	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig() failed,err:%v\n", err)
		return
	}

	//把读取到的配置信息,反序列化到Conf全局变量中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal feiled,err:%v\n", err)
	}

	//监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件修改...\n")
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed,err:%v\n", err)
		}
	})
	return
}
