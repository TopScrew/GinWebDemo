package config

/**
 * @Author: Screw
 * @Date: 2020/10/1 12:25 上午
 * @Desc:
 */

// DevelopConf 开发环境配置
type DevelopConf struct{}


// DataBase 获取数据库配置
func (c *DevelopConf) DataBase() DatabaseConf {

	return DatabaseConf{
		User:     "root",
		Password: "root",
		Host:     "192.168.105.146",
		Port:     3306,
		DBName:   "demo",
	}
}


// DataBase 获取数据库配置
func (c *DevelopConf) Cache() CacheConf {

	return CacheConf{
		Addr:     "127.0.0.1",
		Password: "666666",
		DB:     "haha",

	}
}