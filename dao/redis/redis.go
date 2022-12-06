package redis

import (
	"context"
	goredis "github.com/go-redis/redis/v8"
	"snake-and-ladder/conf"
	"time"
)

var redisClient goredis.UniversalClient

// Init 初始化 redis 连接池，操作示例见 example_test.go
func Init() error {
	c := conf.GetConf().Redis
	redisClient = goredis.NewUniversalClient(&goredis.UniversalOptions{
		Addrs:        c.Addrs,
		DB:           c.DB,
		Password:     c.Password,
		MaxRetries:   c.MaxRetries,
		DialTimeout:  time.Duration(c.DialTimeout) * time.Millisecond,
		ReadTimeout:  time.Duration(c.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(c.WriteTimeout) * time.Millisecond,
		PoolSize:     c.PoolSize,
		MinIdleConns: c.MinIdleConns,
		MaxConnAge:   time.Duration(c.MaxConnAge) * time.Second,
		MasterName:   c.MasterName,
	})

	cmd := redisClient.Ping(context.Background())
	return cmd.Err()
}
