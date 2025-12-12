package init

import (
	"carRider-srv/global"
	"errors"
	"fmt"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

func Nacos() {
	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": []constant.ServerConfig{{IpAddr: "14.103.175.69", Port: 8848}},
		"clientConfigs": []constant.ClientConfig{{NamespaceId: "public"}},
	})
	if err != nil {
		panic(errors.New("连接失败"))
	}
	config, err := client.GetConfig(vo.ConfigParam{
		DataId: "carhub",
		Group:  "DEFAULT_GROUP",
	})
	if err != nil {
		panic(errors.New("获取配置失败"))
	}
	viper.SetConfigType("yaml")
	viper.ReadConfig(strings.NewReader(config))
	err = viper.Unmarshal(&global.AppCfg)
	if err != nil {
		panic(errors.New("配置解析失败"))
	}
	fmt.Println("配置信息: ", global.AppCfg)
}
