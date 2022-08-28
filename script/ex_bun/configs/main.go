package configs

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var (
	Cfg = GetConfig()
)

type config struct {
	DB struct {
		DSN string `mapstructure:"DB_DSN"`
	}
	ARGS_INT int    `mapstructure:"ARGS_INT"`
	ARGS     string `mapstructure:"TEST_ARGS"`
}

func GetConfig() *config {
	_, currentFilePath, _, _ := runtime.Caller(0)
	wd, _ := filepath.Abs(filepath.Dir(currentFilePath))
	// search cfg
	viper.SetConfigType("env")
	viper.AddConfigPath(wd)
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	fmt.Printf("Set AutomaticEnv\n")
	fmt.Printf("TEST_ARGS(env):%+v\n", viper.GetString("TEST_ARGS"))
	fmt.Printf("DB_DSN(env):%+v\n", viper.GetString("DB_DSN"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Can't load config.yaml")
		viper.SetConfigName("default.config")
		err1 := viper.ReadInConfig()
		if err1 != nil {
			log.Println("Can't load example.config.yaml")
		}
	}
	fmt.Printf("viper ReadInConfig\n")
	fmt.Printf("TEST_ARGS(final):%+v\n", viper.GetString("TEST_ARGS"))
	fmt.Printf("DB_DSN(final):%+v\n", viper.GetString("DB_DSN"))

	var cfg config
	err = viper.Unmarshal(&cfg.DB)
	viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable load config into struct, %v", err)
	}
	fmt.Printf("Config:%+v\n", cfg)

	return &cfg
}
