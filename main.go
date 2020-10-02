package main

import (
	"WebDemo/config"
	"WebDemo/router"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	// 获取当前环境
	config.ENV = os.Getenv("ENV")

	// 设置gin模式以及获取服务配置
	var conf config.Conf

	switch config.ENV {
	case "debug":
		gin.SetMode(gin.DebugMode)
		conf = &config.DevelopConf{}
	//case "test":
	//	gin.SetMode(gin.TestMode)
	//	conf = &config.TestingConf{}
	//case "preview":
	//	gin.SetMode(gin.ReleaseMode)
	//	conf = &config.PreviewConf{}
	case "release":
		gin.SetMode(gin.ReleaseMode)
		conf = &config.ReleaseConf{}
	default:
		panic("environment variable error: undefined ENV")
	}

	// 初始化服务依赖
	err := config.Init(conf)
	if err != nil {
		panic("driver init failed: " + err.Error())
	}

	// 初始化Gin路由
	server := router.InitRouter()

	// 启动服务
	server.Run("0.0.0.0:9999")

}
