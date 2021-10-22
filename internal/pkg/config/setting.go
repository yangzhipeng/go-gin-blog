package config

import (
	"time"
)

// 服务器配置
type ServerConfig struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

//数据库配置
type DatabaseConfig struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

// 定义全局变量
var (
	Database *DatabaseConfig
	Server   *ServerConfig
)

// 读取配置到全局变量
func StartSetting() error {
	s, err := NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &Database)
	if err != nil {
		return err
	}

	err = s.ReadSection("Server", &Server)
	if err != nil {
		return err
	}

	return nil
}
