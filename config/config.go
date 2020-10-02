package config

/**
 * @Author: Screw
 * @Date: 2020/9/30 11:47 下午
 * @Desc:
 */


// DatabaseConf 数据库配置属性
type DatabaseConf struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
}

// CacheConf 消息队列配置属性
type CacheConf struct {
	Addr     string
	Password string
	DB       string
}


// Conf 配置功能接口
type Conf interface {

	DataBase() DatabaseConf // 获取数据库配置
	Cache() CacheConf         // 获取缓存数据库配置
}