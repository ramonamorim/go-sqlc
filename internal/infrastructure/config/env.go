package infrastructure

import (
	"cmp"
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	Database     string `mapstructure:"DATABASE"`
	DbConnection string `mapstructure:"DB_CONNECTION"`
}

func LoadEnvConfiguration() *Env {
	env := &Env{}

	envFilepath := cmp.Or(os.Getenv("ENV_FILEPATH"), ".env")

	viper.SetConfigFile(envFilepath)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&env); err != nil {
		panic(err)
	}
	return env
}
