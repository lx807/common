package common

import "github.com/micro/go-micro/config"

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     string `json:"port"`
}

//获取mysql配置
func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	//将配置文件扫描
	config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}
