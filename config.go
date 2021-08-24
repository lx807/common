package common

import (
	"strconv"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-plugins/config/source/consul"
)

//设置配置中心
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		//设置配置中心地址
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		//设置前缀， 不设置为默认前缀
		consul.WithPrefix(prefix),
		//是否移除前缀，true为可以不用前缀直接获取对应配置
		consul.StripPrefix(true),
	)
	// 初始化配置
	config := config.NewConfig()
	//加载配置
	err := config.Load(consulSource)
	return config, err

}
