package config

import (
	"fmt"
	"io"
	"os"

	"github.com/noir017/goapis/pkg/tools"
	"gopkg.in/yaml.v2"
)

type TLS struct {
	Key    string `ymal:"key"`
	CA     string `ymal:"CA"`
	Enable bool   `ymal:"enable"`
}
type Config struct {
	TLS TLS `yaml:"TLS"`
}

func createDefaultConfigFile(configPath string) {
	// 设置默认配置值
	defaultConfig := Config{
		TLS{
			Key:    "",
			CA:     "",
			Enable: false,
		},
	}

	data, err := yaml.Marshal(&defaultConfig)
	if err != nil {
		fmt.Println("Error marshaling default config:", err)
		return
	}

	// 创建或覆盖 YAML 文件
	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
func ReadYml() Config {
	configPath := "config/config.yaml"
	if !tools.IsFile(configPath) {
		createDefaultConfigFile(configPath)
	}
	file, err := os.Open(configPath)
	if err != nil {
		panic("Error opening file:" + err.Error())

	}
	defer file.Close()

	// 读取文件内容
	data, err := io.ReadAll(file)
	if err != nil {
		// fmt.Println()
		panic("Error reading file:" + err.Error())
	}

	var config Config
	// 解析 YAML 数据到结构体
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic("Error parsing YAML:" + err.Error())
	}
	return config
}
