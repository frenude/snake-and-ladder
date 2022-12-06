package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/martian/v3/log"
	"snake-and-ladder/conf"
)

func RunHTTPServer() {
	e := gin.New()
	e.Use(Cors())
	e.Use(gin.Logger())
	// 注册路由
	RouterRegister(e)
	config := conf.GetConf()
	listen := fmt.Sprintf("%s:%d", config.HTTP.Host, config.HTTP.Port)
	log.Infof("http server serving on %s", listen)
	err := e.Run(listen)
	if err != nil {
		log.Errorf("http server start failed: %v", err)
	}
}
