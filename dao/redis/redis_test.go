package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	err := rdb.Set(ctx, "key", 10, time.Hour).Err()
	if err != nil {
		return
	}
	// 直接执行命令获取值
	value := rdb.Get(ctx, "key").Val()
	fmt.Println(value)
	// 执行命令获取结果
	val, err := rdb.Get(ctx, "key").Result()
	fmt.Println(val, err)

}
