package parseini

import (
	"fmt"
	"testing"
)

func TestLoadIni(t *testing.T) {
	//f := os.Open("./config.ini")
	var cfg Config
	err := loadIni("./config.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed err:%v \n", err)
		return
	}
	fmt.Println(cfg)
	fmt.Println(cfg.MysqlConfig)
	fmt.Println(cfg.RedisConfig)
	fmt.Println(cfg.RedisConfig.Database)
}
