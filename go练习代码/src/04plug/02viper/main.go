package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func Initviper() {
	//viper.AddConfigPath("./")
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file change", e.Name)
	})
}

func main() {
	//viper.SetDefault("ContentDir", "aaa")
	//ContentDir := viper.Get("ContentDir")
	//fmt.Println(ContentDir)
	Initviper()
	//viper.Set("name","okk")
	//viper.WriteConfig()
	//viper.AutomaticEnv()
	//viper.BindEnv("GOPATH")
	os.Setenv("GO_PATH","1111")
	viper.SetEnvPrefix("GO")

	fmt.Println(viper.Get("GOPATH"))
	fmt.Println(viper.Get("GO_PATH"))
	//r := gin.Default()
	//r.GET("/hello", func(c *gin.Context) {
	//	c.String(http.StatusOK, viper.GetString("version"))
	//})
	//r.Run(viper.GetString("addr"))
}
