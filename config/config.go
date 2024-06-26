package config

import "github.com/spf13/viper"

// Config represents the environment variables.
type Config struct {
	SECRETKEY    string `mapstructure:"JWTSECRET"`
	Host         string `mapstructure:"HOST"`
	User         string `mapstructure:"USER"`
	Password     string `mapstructure:"PASSWORD"`
	Database     string `mapstructure:"DBNAME"`
	Port         string `mapstructure:"PORT"`
	Sslmode      string `mapstructure:"SSL"`
	GrpcPort     string `mapstructure:"GRPCPORT"`
	ProductPort  string `mapstructure:"GRPCPRODUCTPORT"`
	ChatPort     string `mapstructure:"GRPCCHATPORT"`
	SID          string `mapstructure:"SID"`
	TOKEN        string `mapstructure:"TOKEN"`
	SERVICETOKEN string `mapstructure:"SERVICETOKEN"`
	PHONE        string `mapstructure:"MOBILE"`
	REDISHOST    string `mapstructure:"REDISHOST"`
}

// LoadConfig will load the environment variable to access.
func LoadConfig() *Config {
	var config Config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.Unmarshal(&config)
	return &config
}
