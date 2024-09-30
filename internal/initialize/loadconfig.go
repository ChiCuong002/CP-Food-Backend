package initialize

import (
	"fmt"
	"food-recipes-backend/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config") //path to config
	viper.SetConfigName("local") // file name
	viper.SetConfigType("yaml") // file extension

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration: %w", err))
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("failed to unmarshal configuration: %w", err))
	}
}