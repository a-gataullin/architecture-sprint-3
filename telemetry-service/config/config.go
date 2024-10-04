package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	InfluxDBUrl  string `mapstructure:"INFLUXDB_URL"`
	InfluxDBName string `mapstructure:"INFLUXDB_NAME"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&config)
	return
}
