package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var rdb *redis.Client

// *redis.Client
// redis.NewClient()
// redis.Options{}
// rdb.Ping().Result()

func NewClient() {
	rdb = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "123456",
		DB:           0,
		PoolSize:     50, // 连接池最大连接数，默认cpu * 10
		MinIdleConns: 10, // 连接池最小空闲连接数
		// PoolTimeout: 10  // 取链接后，该链接最多可被使用多长时间
		IdleTimeout:        -1,              // 多久关闭一个空闲链接，默认5min，-1表示不启用
		IdleCheckFrequency: 5 * time.Second, // 多久检测一次空闲连接
	})
	pong, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("ping error...")
		return
	}
	fmt.Println(pong)
}

func main() {
	NewClient()
}
