package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var RootCmd = &cobra.Command{
	Use:           "TEST",
	Short:         "CLI TEST",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default ./config.json )")
	RootCmd.PersistentFlags().StringP("url", "", "example.com", "endpoint")

	viper.BindPFlag("url", RootCmd.PersistentFlags().Lookup("url"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".hoge")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.ReadInConfig()
}
