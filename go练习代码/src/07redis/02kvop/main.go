package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

// 基本的键值操作
// rdb.Set(k, v, time).Err()
// rdb.Get(k).Result()
// rdb.GetSet(k, v).Result()	// 设置一个key的值，并返回这个key的旧值！！！
// redis.Nil
// rdb.SetNX().Err()	// 如果 key 不存在，则设置这个key的值
// rdb.MGet().Result()		// 查询
// rdb.Incr().Result()	// 针对一个key的数值进行递增1操作
// rdb.IncrBy().Result()	// 可以指定每次递增多少
// rdb.Desc().Result()
// rdb.DescBy().Result()
// rdb.Del()		// 删除，支持多个 key

var rdb *redis.Client

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

func GetSetOperator() {
	// 设置key的值，0 表示不会过期
	// err := rdb.Set("key", "value", 0).Err()
	err := rdb.Set("key", "value", time.Second*2).Err()
	if err != nil {
		panic(err)
	}

	// 查询 key 的值
	val, err := rdb.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key：", val)
	time.Sleep(1 * time.Second)

	// 查询key的值，已过期
	val2, err := rdb.Get("key").Result()
	if err == redis.Nil { // 查不到的错误 redis.Nil
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key", val2)
	}

	//  设置新值，返回旧值
	fmt.Println("GetSet")
	oldVal, err := rdb.GetSet("key", "new value").Result()
	if err != redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key", "old value", oldVal)
	}
	fmt.Println("key", "old value", oldVal)
}

func SetNx() {
	fmt.Println("-----------------------NX")
	// 如果 key 不存在，则设置这个key的值，0表示不会过期
	err := rdb.SetNX("keynx", "valuenx", 0).Err()
	if err != nil {
		panic(err)
	}
	v, err := rdb.Get("keynx").Result()
	fmt.Println(v)
	err = rdb.SetNX("keynx", "valuenx2", 0).Err()
	if err != nil {
		panic(err)
	}

	v, err = rdb.Get("keynx").Result() // valuenx
	fmt.Println(v)
}

func MGet() {
	fmt.Println("-------------------", "MGet")
	// 获取多个 key 的值，返回数组
	vals, err := rdb.MGet("key1", "key2", "key3", "key", "keynx").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vals)
	fmt.Println(vals[0])
	if vals[0] == nil {
		fmt.Println("vals[0] is nil")
	}
	fmt.Println(vals[4].(string))
}

func Incr() {
	val, err := rdb.Incr("incr").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("incr最新值", val)
	val2, err := rdb.IncrBy("incrby", 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("incrby最新值", val2)
	fmt.Println("incr的值转Int")
	v, e := rdb.Get("incr").Result()
	if e == redis.Nil {
		fmt.Println("incr not exist...")
	} else if e != nil {
		panic(err)
	} else {
		fmt.Println(v)
		i, _ := strconv.ParseInt(v, 10, 0)
		fmt.Println(i)
	}
}

// Decr, DecrBy 减法
func DecrBy() {
	val, err := rdb.DecrBy("incr", 10000).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}

// Del 删除
func Del() {
	rdb.Del("incr") // 删除 一个 Key		// 不写 Result，不接收了，太累了...
	rdb.Del("key", "incrby")
}

// 设置 key 的过期时间，单位秒
func Expire() {
	rdb.Expire("keynx", 3*time.Second) // 不写 Result，不接收了，太累了...
}

func main() {
	NewClient()
	GetSetOperator()
	SetNx()
	MGet()
	Incr()
	DecrBy()
	Del()
	Expire()
}
