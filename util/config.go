package util

import "github.com/spf13/viper"

type Config struct {
	DB_DB_HOST  string `mapstructure:"DB_DB_HOST"`
	DB_USER     string `mapstructure:"DB_USER"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	DB_DBNAME   string `mapstructure:"DB_DBNAME"`
	DB_PORT     string `mapstructure:"DB_PORT"`
	DB_SSLMODE  string `mapstructure:"DB_SSLMODE"`

	API_KEY string `mapstructure:"API_KEY"`
	URL     string `mapstructure:"URL"`

	SALT        string `mapstructure:"SALT"`
	SIGNING_KEY string `mapstructure:"SIGNING_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
