package dao

import (
	"fmt"
	"github.com/google/martian/v3/log"
	"snake-and-ladder/dao/db"
	"snake-and-ladder/dao/redis"
)

func Init() error {
	err := db.Init()
	if err != nil {
		return fmt.Errorf("db init: %v", err)
	}
	err = redis.Init()
	if err != nil {
		// todo 申请redis, 打开 return 语句
		log.Errorf("redis init: %v", err)
	}
	return nil
}
