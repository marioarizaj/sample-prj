package utilities

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type IConfigUtil interface {
	GetConfig(key string) string
	GetEnv() string
}

type ConfigUtil struct {
	Viper *viper.Viper
}

func NewConfigUtil() *ConfigUtil {
	// read config
	var env string
	if os.Getenv("GO_ENV") != "" {
		env = os.Getenv("GO_ENV")
	} else {
		env = "development"
	}
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "TEST":
			env = "test"
		case "PROD":
			env = "prod"
		case "REG":
			env = "reg"
		}
	}
	r := ConfigUtil{}
	r.Viper = viper.New()
	r.Viper.SetConfigName(env)
	r.Viper.AddConfigPath("./")
	r.Viper.AddConfigPath("./config/")
	r.Viper.AddConfigPath(".")
	r.Viper.SetConfigType("json")
	r.Viper.AddConfigPath("/go/src/bitbucket.org/ProNovate/quoter-services2-go/config/")
	err := r.Viper.ReadInConfig() // Find and read the config file
	if err != nil {
		// Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return &r
}

func (r ConfigUtil) GetConfig(key string) string {
	return r.Viper.GetString(key)
}

func (r ConfigUtil) GetEnv() string {
	var env string
	if os.Getenv("GO_ENV") != "" {
		env = os.Getenv("GO_ENV")
	} else {
		env = "development"
	}
	return env
}
