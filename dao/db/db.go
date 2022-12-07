package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/url"
	"snake-and-ladder/conf"
	"snake-and-ladder/model/po"
	"time"
)

var dbClient *gorm.DB

// Init 初始化 db
func Init() error {
	log.Println("start init db")

	var err error
	dbClient, err = gorm.Open(dialector(), &gorm.Config{})
	if err != nil {
		return err
	}

	// 获取驱动实例，设置连接池参数
	sqlDB, err := dbClient.DB()
	if err != nil {
		return err
	}
	config := conf.GetConf()
	sqlDB.SetMaxIdleConns(config.DB.MaxIdle)
	sqlDB.SetMaxOpenConns(config.DB.MaxConn)
	sqlDB.SetConnMaxLifetime(time.Duration(config.DB.ConnLifeTime) * time.Second)
	if !dbClient.Migrator().HasTable(&po.Board{}) {
		dbClient.Migrator().CreateTable(&po.Board{})
	}
	if !dbClient.Migrator().HasTable(&po.Step{}) {
		dbClient.Migrator().CreateTable(&po.Step{})
	}
	log.Println("db init success")
	return nil
}

func dialector() gorm.Dialector {
	c := conf.GetConf().DB

	if c.IsPostgres {
		return postgres.Open(
			fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
				c.Host, c.User, c.Password, c.DataBase, c.Port, url.QueryEscape(c.TimeZone)))
	}
	return mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=%dms&readTimeout=%dms&writeTimeout=%dms&parseTime=%t&loc=%s",
		c.User, c.Password, c.Host, c.Port, c.DataBase, c.Charset, c.ConnTimeout, c.ReadTimeout,
		c.WriteTimeout, c.ParseTime, url.QueryEscape(c.TimeZone)))
}
