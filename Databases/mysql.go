package Databases

/**
 * @Author: Screw
 * @Date: 2020/9/30 9:57 下午
 * @Desc:
 */
import (
	"WebDemo/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func initDB() (err error) {

	//格式  账号:密码@(ip:端口)/数据库?charset=utf8mb4编码等参数
	db, err = gorm.Open("mysql", "root:root@(192.168.105.129:3306)/gorm_study?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic("连接数据库失败:" + err.Error())
	} else {
		fmt.Println("数据库连接成功~~~~~~")
	}

	defer db.Close()

	//SetMaxOpenConns用于设置最大打开的连接数
	db.DB().SetMaxIdleConns(10)
	//SetMaxIdleConns用于设置闲置的连接数
	db.DB().SetMaxOpenConns(100)

	// 启用Logger，显示详细日志
	db.LogMode(true)

	// 自动迁移创建新表，把数据体和数据表进行对应
	db.AutoMigrate(&models.User{})

	return
}
