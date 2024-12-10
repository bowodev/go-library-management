package config

import "time"

type Config struct {
	Env     string `mapstructure:"env"`
	AppName string `mapstructure:"appName"`
	Port    string `mapstructure:"port"`
	DB      DB     `mapstructure:"db"`
	Redis   Redis  `mapstructure:"redis"`
}

type DB struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"dbName"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Redis struct {
	Address         string    `mapstructure:"address"`
	DB              int       `mapstructure:"db"`
	BookCache       CacheData `mapstructure:"bookCache"`
	AuthorCache     CacheData `mapstructure:"authorCache"`
	BookAuthorCache CacheData `mapstructure:"bookAuthorCache"`
}

type CacheData struct {
	Key string        `mapstructure:"key"`
	TTL time.Duration `mapstructure:"ttl"`
}
