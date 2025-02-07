package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
	Database struct {
		Path string `json:"path"`
	} `json:"database"`
	JWT struct {
		Secret string `json:"secret"`
		Expire int    `json:"expire"` // token过期时间（小时）
	} `json:"jwt"`
}

var GlobalConfig Config

// LoadConfig 从文件加载配置
func LoadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&GlobalConfig)
}
