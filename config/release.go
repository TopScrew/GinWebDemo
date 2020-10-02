package config

/**
 * @Author: Screw
 * @Date: 2020/10/2 12:45 下午
 * @Desc:
 */
// DevelopConf 开发环境配置
type ReleaseConf struct{}


// DataBase 获取数据库配置
func (c *ReleaseConf) DataBase() DatabaseConf {

	return DatabaseConf{
		User:     "root",
		Password: "root",
		Host:     "192.168.105.145",
		Port:     3306,
		DBName:   "demo",
	}
}


// DataBase 获取数据库配置
func (c *ReleaseConf) Cache() CacheConf {

	return CacheConf{
		Addr:     "127.0.0.1",
		Password: "666666",
		DB:     "haha",

	}
}