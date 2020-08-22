package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Env *config

type config struct {
	DSN string
}

func init() {
	var v = viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config file open fail"))
	}

	if err := v.Unmarshal(&Env); err != nil {
		panic(fmt.Errorf("config struct parse err"))
	}
	fmt.Println(Env)
}
