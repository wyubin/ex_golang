package sqlkit

import (
	"ex_golang/structure/pkg/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (c *PGConfig) AddFlags(cmd *cobra.Command, namespace string) {
	urlKey := namespace + ".url"
	debugKey := namespace + ".debug"
	poolSizeKey := namespace + ".poolsize"

	cmd.Flags().StringVar(&c.URL, urlKey, "postgresql://postgres:postgres@localhost?sslmode=disable", "postgresql url")
	cmd.Flags().BoolVar(&c.Debug, debugKey, false, "postgresql debug")
	cmd.Flags().IntVar(&c.PoolSize, poolSizeKey, 5, "postgresql pool size")

	viper.BindPFlag(urlKey, cmd.Flags().Lookup(urlKey))
	viper.BindPFlag(debugKey, cmd.Flags().Lookup(debugKey))
	viper.BindPFlag(poolSizeKey, cmd.Flags().Lookup(poolSizeKey))

	viper.BindEnv(urlKey, utils.ToEnvKey(urlKey))
	viper.BindEnv(debugKey, utils.ToEnvKey(debugKey))
	viper.BindEnv(poolSizeKey, utils.ToEnvKey(poolSizeKey))

	c.URL = viper.GetString(urlKey)
	c.Debug = viper.GetBool(debugKey)
	c.PoolSize = viper.GetInt(poolSizeKey)
}
