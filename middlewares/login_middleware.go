package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: Screw
 * @Date: 2020/10/2 4:37 下午
 * @Desc:
 */
func LoginHandlerFunc(c *gin.Context)  {

	fmt.Println("判断是否登录")

	c.Set("username","screw")

	//执行程序的后续请求
	c.Next()
	//阻止程序的后续请求
	//c.Abort()
}