package config

import (
	"WebDemo/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

/**
 * @Author: Screw
 * @Date: 2020/10/2 1:01 下午
 * @Desc:
 */

// 全局变量
var (
	DB        *gorm.DB
	ENV       string
)


// 创建MySQL客户端
func NewMysqlClient(conf DatabaseConf) (err error) {
	openArgs := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DBName)
	db, err := gorm.Open("mysql", openArgs)
	if err != nil {
		return
	}
	err = db.DB().Ping()
	if err != nil {
		return
	}

	if os.Getenv("ENV") == "debug" || os.Getenv("ENV") == "test" { // 开发和测试环境自动执行
		db.AutoMigrate( // Migrate the schema
			&models.User{},
			//&model.AccessControlList{},   //迁移表
		)
		db.LogMode(true)
	}
	//SetMaxOpenConns用于设置最大打开的连接数
	db.DB().SetMaxIdleConns(10)
	//SetMaxIdleConns用于设置闲置的连接数
	db.DB().SetMaxOpenConns(50)

	// 启用Logger，显示详细日志
	db.LogMode(true)

	DB = db
	return
}


func Init(conf Conf) (err error)  {

	err = NewMysqlClient(conf.DataBase())

	return
}