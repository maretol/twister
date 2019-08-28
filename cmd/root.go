package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Hoge string
var cfgFile string
var RootCmd = &cobra.Command{
	Use:           "twister",
	Short:         "Twister is web differ",
	Long:          "Hello! Welcome to twister which is web differ.\nfirst, you run the command 'twister init'",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "twisterconf.json", "config file")

	viper.BindPFlag("url", RootCmd.PersistentFlags().Lookup("url"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName("twisterconf")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.twister")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		Config = ConfigBase{}
		return
	}

	Config.Urls = viper.GetStringSlice("urls")
	Config.Paths = viper.GetStringSlice("paths")
	Config.Check.Body = viper.GetBool("check.body")
	Config.Check.Header = viper.GetBool("check.header")
	Config.Check.StatusCode = viper.GetBool("checker.statusCode")
	IsConfigExist = true
}

var IsConfigExist = false
var Config ConfigBase

type ConfigBase struct {
	Urls  []string `json:"urls"`
	Paths []string `json:"paths"`
	Check checker  `json:"check"`
}

type checker struct {
	Body       bool `json:"body"`
	StatusCode bool `json:"statusCode"`
	Header     bool `json:"header"`
}
