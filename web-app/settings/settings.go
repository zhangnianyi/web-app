package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigFile("config.yaml") // 指定配置文件
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath(".")     // 指定查找配置文件的路径
	err = viper.ReadInConfig()        // 读取配置信息
	if err != nil {                    // 读取配置信息失败
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		return
	}
	viper.WatchConfig()
	// 注意！！！配置文件发生变化后要同步到全局变量Conf
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("夭寿啦~配置文件被人修改啦...")

	})
	return
}
