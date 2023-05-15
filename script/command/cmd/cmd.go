package cmd

import (
	"fmt"
	"os"

	"example.com/command/utils/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

var (
	cfgFile string
	ckDebug bool

	democtlCmd = &cobra.Command{
		Use:           "democtl",
		Short:         "democtl – command-line practice",
		Long:          ``,
		Version:       "0.1.0",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func Execute() error {
	return democtlCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	democtlCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.democtl.yaml)")
	democtlCmd.PersistentFlags().BoolVar(&ckDebug, "debug", false, "show debug log if needed")
	viper.BindPFlag("config", democtlCmd.PersistentFlags().Lookup("config"))

	democtlCmd.AddCommand(addCmd)
	democtlCmd.AddCommand(lsCmd)
}

func initConfig() {
	// init logger
	var logLevel zapcore.Level = zapcore.InfoLevel
	if ckDebug {
		logLevel = zapcore.DebugLevel
	}
	log.InitLogger(logLevel)
	// load yaml if has
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigFile(".democtl")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Logger.Info(fmt.Sprint("Config file used for democtl: ", viper.ConfigFileUsed()))
	}
}
