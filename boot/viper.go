package boot

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	g "juejin/app/global"
	"os"
)

// 优先级：命令行参数>环境变量>默认
const (
	configEnv  = "JUEJIN_CONFIG_PATH"
	configFile = "manifest/config/config.yaml"
)

func ViperSetup(path ...string) {
	var configPath string
	if len(path) != 0 {
		configPath = path[0]
	} else {
		flag.StringVar(&configPath, "c", "", "set config file path")
		flag.Parse()
		if configPath == "" {
			if configPath = os.Getenv(configEnv); configPath != "" {

			} else {
				configPath = configFile
			}
		}

	}
	fmt.Printf("get configPath:%s", configPath)
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("get config file failed,err:%v", err))
	}
	if err := v.Unmarshal(&g.Config); err != nil {
		panic(fmt.Errorf("unmarshal config failed, err:%v", err))
	}

}
