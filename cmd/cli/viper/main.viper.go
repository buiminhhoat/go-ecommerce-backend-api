package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}

	fmt.Println("Server.Port::", viper.GetInt("server.port"))
	fmt.Println("Key::", viper.GetString("security.jwt.key"))

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration %w", err)
	}

	fmt.Println("Config port::", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("databases user: %s, password: %s\n", db.User, db.Password)
	}
}
