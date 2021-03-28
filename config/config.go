package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const active_name string = "active"

func GetDataSourceConfig() interface{} {
	a := Active()
	config := viper.New()
	config.AddConfigPath("resource/config") //设置读取的文件路径
	config.SetConfigName(configName(a))     //设置读取的文件名
	config.SetConfigType("yaml")
	config.ReadInConfig()
	return config.Get("datasource")
}

func Active() string {
	//获取项目的执行路径
	config := viper.New()
	config.AddConfigPath("resource/config") //设置读取的文件路径
	config.SetConfigName(configName(""))    //设置读取的文件名
	config.SetConfigType("yaml")            //设置文件的类型
	//尝试进行配置读取
	config.ReadInConfig()
	a := config.Get(active_name)
	res := fmt.Sprintf("%s", a)
	return res
}

func configName(active string) string {
	if len(active) == 0 {
		return ""
	}
	return fmt.Sprintf("application-%s.yml", active)
}
