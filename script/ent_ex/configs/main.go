package configs

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var (
	Cfg = GetConfig()
)

type config struct {
	Gin commonCfg
}

type commonCfg struct {
	Port int
}

func GetConfig() *config {
	_, currentFilePath, _, _ := runtime.Caller(0)
	wd, _ := filepath.Abs(filepath.Dir(currentFilePath))
	// search cfg
	viper.SetConfigType("yaml")
	viper.AddConfigPath(wd)
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Can't load config.yaml")
		viper.SetConfigName("example.config")
		err1 := viper.ReadInConfig()
		if err1 != nil {
			log.Fatal("Can't load example.config.yaml")
		}
	}
	var cfg config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("unable load config into struct, %v", err)
	}
	return &cfg
}
