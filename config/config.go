package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// AppConfig 存储应用程序配置
type AppConfig struct {
	Autostart bool `json:"autostart"`
}

const configFileName = "config.json"

// GetConfigPath 返回配置文件路径
func GetConfigPath() string {
	execPath, err := os.Executable()
	if err != nil {
		return configFileName
	}
	return filepath.Join(filepath.Dir(execPath), configFileName)
}

// LoadConfig 从文件加载配置
func LoadConfig() *AppConfig {
	conf := &AppConfig{
		Autostart: false, // 默认不自启
	}

	data, err := os.ReadFile(GetConfigPath())
	if err != nil {
		return conf
	}

	_ = json.Unmarshal(data, conf)
	return conf
}

// SaveConfig 保存配置到文件
func (c *AppConfig) SaveConfig() error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(GetConfigPath(), data, 0644)
}
