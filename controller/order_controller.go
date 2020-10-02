package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: Screw
 * @Date: 2020/9/30 6:35 下午
 * @Desc:
 */

func Sayhello(context *gin.Context)  {

	context.JSON(200,gin.H{
		"message": "hello gin~~~GET~",
	})

	name, _ := context.Get("username")

	fmt.Println("~~~~~~~~~~~",name)
}