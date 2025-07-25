package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"sync"
)

// 定义全局配置结构
var (
	config *Configuration
	once   sync.Once
)

// Configuration 结构体对应配置文件内容
type Configuration struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	
	Log struct {
		Path  string `yaml:"path"`
		Level string `yaml:"level"`
	} `yaml:"log"`

	MySQL struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"mysql"`

	Auth struct {
		JwtSecret   string `yaml:"jwt_secret"`
		TokenExpiry int    `yaml:"token_expiry"`
	} `yaml:"auth"`
}

// InitConfig 初始化配置（单例模式）
func InitConfig(configPath string) {
	once.Do(func() {
		config = &Configuration{}
		if err := loadConfig(configPath); err != nil {
			log.Fatalf("yaml配置加载失败: %v", err)
		}
	})
}

// GetConfig 获取全局配置实例
func GetConfig() *Configuration {
	return config
}

// 加载并解析配置文件
func loadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("加载配置文件失败: %w", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return fmt.Errorf("解析配置文件失败: %w", err)
	}
	if config.Server.Port == "" {
		config.Server.Port = "8080"
	}
	if config.MySQL.Port == 0 {
		config.MySQL.Port = 3306
	}
	return nil
}
