package configs

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

var (
	Cfg = GetConfig()
)

type config struct {
	DSN   string `mapstructure:"db_dsn"`
	Debug bool   `mapstructure:"db_debug"`
}

type exCfg struct {
	ARGS_INT int `mapstructure:"ARGS_INT"`
}

func GetConfig() *config {
	_, currentFilePath, _, _ := runtime.Caller(0)
	wd, _ := filepath.Abs(filepath.Dir(currentFilePath))
	// search cfg
	//viper.SetConfigType("yaml")
	viper.AddConfigPath(wd)
	viper.SetConfigName("default_config")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("DB_", "SQLITE_"))
	// viper.SetEnvPrefix("sqlite")
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
	var cfg1 exCfg
	for _, cfgTmp := range []interface{}{&cfg, &cfg1} {
		err = viper.Unmarshal(cfgTmp)
		if err != nil {
			log.Fatalf("unable load config into struct, %v", err)
		}
		fmt.Printf("Config:%+v\n", cfgTmp)
	}

	return &cfg
}
