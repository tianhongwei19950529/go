package main

import (
	"05demo/dao/mysql"
	"05demo/routes"
	"05demo/settings"
	"05demo/utils"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	//加载配置文件
	if err := settings.Init(); err != nil {
		fmt.Println(err.Error())
		return
	}
	logger.Init()
	zap.L().Info("日志初始化完成....")
	defer zap.L().Sync()
	err := mysql.Init()
	defer mysql.Close()
	if err != nil {
		zap.L().Error("db connect fail...", zap.Error(err))
	}
	r := routes.SetUp()
	r.Run(viper.GetString("app.port"))

}
