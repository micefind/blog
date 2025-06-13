package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App   AppConfig   `yaml:"app"`
	Mysql MysqlConfig `yaml:"mysql"`
	Jwt   JwtConfig   `yaml:"jwt"`
}

type AppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Port    int    `yaml:"port"`
}

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
}

type JwtConfig struct {
	Secret string `yaml:"secret"`
}

var (
	App   AppConfig
	Mysql MysqlConfig
	Jwt   JwtConfig
)

func LoadConfig(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var cfg Config
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return err
	}
	App = cfg.App
	Mysql = cfg.Mysql
	Jwt = cfg.Jwt

	return nil
}
