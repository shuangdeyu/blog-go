package conf

import (
	"io/ioutil"
	"github.com/go-yaml/yaml"
)

type Configuration struct {
	DataBaseUrl   string `yaml:"DataBaseUrl"`
	Port          string `yaml:"Port"`
	RedisHost     string `yaml:"RedisHost"`
	RedisPassword string `yaml:"RedisPassword"`
}

var configuration *Configuration // 相当于声明一个指针变量

// 加载配置文件
func LoadConfiguration(path string) error {
	// 读取文件
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var config Configuration
	// 解析文件数据，并存入结构体
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	// 将解析后的数据赋值给全局变量
	configuration = &config
	return err
}

// 获取配置
func GetConfiguration() *Configuration {
	return configuration
}