package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// Config Gitlab CLI Config
type Config struct {
	Token     string
	Host      string
	WorkSpace string
}

var configName = "config.json"
var configPath = "/usr/local/share/cinema/gitlab"

func init() {

}

func fetchConfig() *Config {
	//查找本地配置文件，不存在则初始化
	configFile := path.Join(configPath, configName)
	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		return initDefault()
	}

	// 读取配置信息
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// 配置信息模型转换
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &config
}

// 初始化默认配置
func initDefault() *Config {
	config := Config{}
	return &config
}

// UpdateConfig 更新配置
func UpdateConfig(config *Config) error {
	if config == nil {
		//TODO: log config is nil
		return nil
	}
	infoJSON, _ := json.Marshal(config)
	filePath := path.Join(config.WorkSpace, configName)
	write(infoJSON, filePath)
	return nil
}

func write(json []byte, filePath string) {
	if json != nil {
		dirPath := path.Dir(filePath)
		mkErr := os.MkdirAll(dirPath, os.ModePerm)
		if mkErr != nil {
			log.Fatal(mkErr)
			return
		}
		writeErr := ioutil.WriteFile(filePath, json, os.ModePerm)
		if writeErr != nil {
			log.Fatal(writeErr)
			return
		}
	}
}
