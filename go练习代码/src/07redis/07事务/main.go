package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"reflect"
	"time"
)

// pipe := rdb.TxPipeLine()
// pipe.xxx
// pipe.Exec()

var rdb *redis.Client

// HSetNX

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

// pipeline 方式
func main1() {
	pipe := rdb.TxPipeline()
	incr := pipe.Incr("tx_pipeline_counter")
	fmt.Println(incr)
	pipe.Expire("tx_pipeline_counter", time.Hour)
	exec, err := pipe.Exec()
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(exec[1]), exec[1])
}

func main2() {
	fn := func(tx *redis.Tx) error {
		v := tx.Get("tx_pipeline_counter").Val()
		// 处理业务....
		fmt.Println(v)
		// ...
		// 如果watch的key没有改变的话，pipelined 函数才会调用成功
		_, err := tx.Pipelined(func(pipeliner redis.Pipeliner) error {
			// 可以在这里给 watch的 key 赋新值
			_, err := pipeliner.Incr("tx_pipeline_counter").Result()
			if err != nil {
				return err
			}
			return nil
		})
		return err
	}

	err := rdb.Watch(fn, "tx_pipeline_counter") // NOTE:支持多个key
	if err != nil {
		return
	}
}

func main() {
	NewClient()
	main1()
	main2()
}
