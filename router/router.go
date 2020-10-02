package router

/**
 * @Author: Screw
 * @Date: 2020/9/30 4:29 下午
 * @Desc:
 */

import (
	"WebDemo/controller"
	"WebDemo/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	engine := gin.Default()
	//engine := gin.New()

	v1 := engine.Group("/api/v1")

	//给v1注册一个全局的登录中间件
	v1.Use(middlewares.LoginHandlerFunc)

	{

		order := v1.Group("/order")

		order.GET("/list", controller.Sayhello)
		order.GET("/top", controller.Sayhello)

		user := v1.Group("/user")

		user.GET("/list", controller.Sayhello)
		user.GET("/top", controller.Sayhello)
	}

	return engine
}
