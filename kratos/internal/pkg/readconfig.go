package pkg

import (
	"fmt"
	"github.com/spf13/viper"
)

/**
 * 通用的读取配置文件的方法，传入路径和文件名以及类型，返回一个Viper的指针
 */
func readFile(filePath, fileName, configType string) (config *viper.Viper, err error) {
	config = viper.New()
	config.AddConfigPath(filePath)
	config.SetConfigName(fileName)
	config.SetConfigType(configType)
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Not found config file")
		} else {
			fmt.Printf("read config file error, %v \n", err)
		}
	}
	return
}

// ReadIniConfig 读取ini类型配置文件
func ReadIniConfig(filePath, fileName string) (*viper.Viper, error) {
	return readFile(filePath, fileName, "ini")
}

// ReadJsonConfig 读取json类型配置文件
func ReadJsonConfig(filePath, fileName string) (*viper.Viper, error) {
	return readFile(filePath, fileName, "json")
}

// ReadJsonConfigAndTransToStruct 读取json配置文件并转成制定的结构体
func ReadJsonConfigAndTransToStruct(filePath, fileName string, value any) error {
	config, err := ReadJsonConfig(filePath, fileName)
	if err != nil {
		return err
	}
	return config.Unmarshal(value)
}

// ReadYamlConfig 读取yaml类型配置文件
func ReadYamlConfig(filePath, fileName string) (*viper.Viper, error) {
	return readFile(filePath, fileName, "yaml")
}
