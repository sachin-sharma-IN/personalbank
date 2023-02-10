package util

import "github.com/spf13/viper"

// Config stroes all config of the appl.
// The values are read by viper from a config file or env. variable.
// Viper uses Mapstructure package under the hood for unmarshalling values so we'll use mapstructure tag.
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// viper.AutomaticEnv() overrides config values with env values if matching env var is found in config file.
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
