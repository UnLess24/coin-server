package config

import "github.com/spf13/viper"

type Config struct {
	Host        string
	APIKey      string
	LastListUrl string
}

func MustRead() Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return Config{
		Host:        viper.GetString("host"),
		APIKey:      viper.GetString("api_key"),
		LastListUrl: viper.GetString("last_list_url"),
	}
}
