package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server         Server
	CoinMarketCap  CoinMarketCap
	TickerDuration time.Duration
	RequestTimeout time.Duration
	Redis          Redis
}

type CoinMarketCap struct {
	Host        string
	APIKey      string
	LastListUrl string
}

type Redis struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type Server struct {
	Host string
	Port string
	Type string
}

func MustRead() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return &Config{
		Server: Server{
			Host: viper.GetString("server.host"),
			Port: viper.GetString("server.port"),
			Type: viper.GetString("server.type"),
		},
		CoinMarketCap: CoinMarketCap{
			Host:        viper.GetString("coinmarketcap.host"),
			APIKey:      viper.GetString("coinmarketcap.api_key"),
			LastListUrl: viper.GetString("coinmarketcap.last_list_url"),
		},
		TickerDuration: viper.GetDuration("ticker_duration"),
		RequestTimeout: viper.GetDuration("request_timeout"),
		Redis: Redis{
			Host:     viper.GetString("redis.host"),
			Port:     viper.GetString("redis.port"),
			Password: viper.GetString("redis.password"),
			DB:       viper.GetInt("redis.db"),
		},
	}
}
