package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

// rdb.HSet()
// rdb.HGet()
// rdb.HGetAll()
// rdb.HIncrBy()  hash 指定字段一次性累加
// rdb.HKeys()	返回该hash所有字段名
// rdb.HLen()	根据 key，查询hash的字段数量
// rdb.HMSet()	批量设置hash字段值
// rdb.HMGet()	批量查询hash字段值
// rdb.HSetNX()	如果字段field不存在，则设置 hash 字段值
// rdb.HDel()	// 根据key和字段名，删除hash字段. 支持批量删除

var rdb *redis.Client

func HSetAndHGetAndAll() {
	funcName, _, _, ok := runtime.Caller(0)
	if ok {
		fmt.Println("func:", runtime.FuncForPC(funcName).Name())
	}
	fmt.Println()
	// 可设置字符串和int型，统一字符串好一点
	_ = rdb.HSet("userw", "username_1", "tizi365").Err()
	err := rdb.HSet("userw", "username_2", 123).Err()
	if err != nil {
		panic(err)
	}

	username, err := rdb.HGet("userw", "username_2").Result()
	if err != nil {
		panic(err)
	}
	parseInt, _ := strconv.ParseInt(username, 10, 64)
	fmt.Println(parseInt + 4)

	data, err := rdb.HGetAll("userw").Result()
	for field, val := range data {
		fmt.Println(field, val)
	}
}

// 根据key和field的值，累加字段的数值
func HIncrBy() {
	funcName, _, _, ok := runtime.Caller(0)
	if ok {
		fmt.Println("func:", runtime.FuncForPC(funcName).Name())
	}
	fmt.Println("HIncrBy")
	count, err := rdb.HIncrBy("user_hash", "count", 4).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}

// 根据 Key 返回所有字段名
func HKeysAndLen() {
	funcName, _, _, ok := runtime.Caller(0)
	if ok {
		fmt.Println("func:", runtime.FuncForPC(funcName).Name())
	}
	keys, err := rdb.HKeys("user_hash").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)
	// size, err := rdb.HLen("user_hash").Result()
	size, err := rdb.HLen("userw").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(size)
}

// HMGet
func HMGetAndHMSet() {
	funcName, _, _, ok := runtime.Caller(0)
	if ok {
		fmt.Println("func:", runtime.FuncForPC(funcName).Name())
	}
	vals, err := rdb.HMGet("user_hash", "count", "user1", "user2", "user3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vals)
	fmt.Println(vals[0])
	fmt.Println(vals[1])
	fmt.Println(reflect.TypeOf(vals))
	data := make(map[string]interface{})
	data["user1"] = 1
	data["user2"] = "tizi"
	err = rdb.HMSet("user_hash", data).Err()
	if err != nil {
		panic(err)
	}
}

func HSetNXAndHDelAndHExists() {
	funcName, _, _, ok := runtime.Caller(0)
	if ok {
		fmt.Println("func:", runtime.FuncForPC(funcName).Name())
	}
	err := rdb.HSetNX("user_hash", "user3", "fuck").Err()
	if err != nil {
		panic(err)
	}
	keys, err := rdb.HKeys("user_hash").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)
	rdb.HDel("user_hash", "user3", "count")
	keys, err = rdb.HKeys("user_hash").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)
	// err = rdb.HExists("user_hash", "count").Err()
	exist, err := rdb.HExists("user_hash", "user1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(exist)
}

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

func main() {
	NewClient()
	HSetAndHGetAndAll()
	HIncrBy()
	HKeysAndLen()
	HMGetAndHMSet()
	HSetNXAndHDelAndHExists()
}
