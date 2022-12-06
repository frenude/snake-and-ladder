package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"snake-and-ladder/application/middlewares"
	"snake-and-ladder/conf"
)

// RouterRegister 注册 http 路由
func RouterRegister(e *gin.Engine) {
	e.GET("/", Hello)
	config := conf.GetConf()
	version := fmt.Sprintf("/api/%v", config.Server.Version)
	//添加路由版本
	vGroup := e.Group(version)
	// 随机棋盘
	vGroup.POST("/randomboard", RandomBoard)
	AGroup := vGroup.Group("/admin")
	AGroup.Use(middlewares.JwtAuthMiddleware())
	// 随机骰子
	AGroup.GET("/randomdice", RandomDice)
}
