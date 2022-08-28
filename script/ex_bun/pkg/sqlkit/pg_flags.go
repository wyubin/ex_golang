package sqlkit

import (
	"strings"

	"example.com/ex_bun/utils/path"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type PGConfig struct {
	URL      string `mapstructure:"pg_url"`
	Debug    bool   `mapstructure:"pg_debug"`
	PoolSize int    `mapstructure:"pg_poolsize"`
}

// 可基於 cobra 加上設定
func (s *PGConfig) AddFlags(cmd *cobra.Command, namespace string) {
	scriptDir := path.DirScript()
	// setup default config
	viper.SetConfigType("yaml")
	viper.AddConfigPath(scriptDir)
	viper.SetConfigName("pg_default")
	// bind env
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("PG_", strings.ToUpper(namespace)+"_"))
	// load config
	viper.ReadInConfig()
	viper.Unmarshal(s)

	cmd.Flags().StringVar(&s.URL, strings.Join(namespace, "url"), s.URL, "postgresql url")
	cmd.Flags().BoolVar(&s.Debug, strings.Join(namespace, "debug"), s.Debug, "postgresql debug")
	cmd.Flags().IntVar(&s.PoolSize, strings.Join(namespace, "poolsize"), s.PoolSize, "postgresql pool size")

}
