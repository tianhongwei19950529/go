package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// set 类型，集合操作
// rdb.SAdd()	// 添加集合元素
// rdb.SCard()	// 获取集合元素个数
// rdb.SIsMember()	// 判断元素是否在集合汇总
// rdb.SMembers()	// 获取集合中所有的元素
// rdb.SRem()		// 删除集合元素
// rdb.SPop, SPopN	// 随机返回集合中的元素，并且删除返回的元素

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

func SAddSCardSIsMemberSMembersSRemSPopSPopN() {
	err := rdb.SAdd("key_set", 100).Err()
	if err != nil {
		panic(err)
	}
	rdb.SAdd("key_set", 100, 200, 300)
	size := rdb.SCard("key_set") // 获取集合元素个数
	fmt.Println("----------")
	fmt.Println(size.Val())
	fmt.Println(size.Name()) // 命令名称
	exist := rdb.SIsMember("key_set", "100").Val()
	fmt.Println(exist)
	exist = rdb.SIsMember("key_set", 100).Val()
	fmt.Println(exist)
	fmt.Println(rdb.SMembers("key_set").Val())
	rdb.SRem("key_set", 100)
	fmt.Println(rdb.SMembers("key_set").Val())
	// 随机返回集合中的一个元素，并删除这个元素
	fmt.Println("-------------SPOP")
	val := rdb.SPop("key_set").Val()
	fmt.Println(val)
	fmt.Println(rdb.SMembers("key_set").Val())
	fmt.Println("-------------SPOPN")
	vals := rdb.SPopN("key_set", 5).Val()
	fmt.Println(vals)
	fmt.Println(rdb.SMembers("key_set").Val())
}

func main() {
	NewClient()
	SAddSCardSIsMemberSMembersSRemSPopSPopN()
}
