package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd はコマンド定義そのもの。これに AddCommand していく
var RootCmd = &cobra.Command{
	Use:           "twister [command]",
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
	rowUrls := viper.Get("urls")
	urls := rowUrls.([]interface{})
	Config.Urls = []Urls{}
	switch urls[0].(type) {
	case string:
		for _, obj := range urls {
			u := Urls{Tag: obj.(string), URL: obj.(string)}
			Config.Urls = append(Config.Urls, u)
		}
		break
	case Urls:
		for _, obj := range urls {
			Config.Urls = append(Config.Urls, obj.(Urls))
		}
		break
	case map[string]interface{}:
		fmt.Printf("here\n")
		for _, obj := range urls {
			u := obj.(map[string]interface{})["url"]
			t := obj.(map[string]interface{})["tag"]
			o := Urls{Tag: t.(string), URL: u.(string)}
			Config.Urls = append(Config.Urls, o)
		}
		break
	default:
		log.Fatal("ERROR! URL format is not correct in config file")
	}
	Config.Paths = viper.GetStringSlice("paths")
	Config.Check.Body = viper.GetBool("check.body")
	Config.Check.Header = viper.GetBool("check.header")
	Config.Check.StatusCode = viper.GetBool("check.statusCode")
	IsConfigExist = true
}

// IsConfigExist はコンフィグファイルが見つけられたかどうか
var IsConfigExist = false

// Config はコンフィグファイルから読んだ情報
var Config ConfigBase

// ConfigBase はコンフィグ定義をしている構造体
type ConfigBase struct {
	Urls  []Urls   `json:"urls"`
	Paths []string `json:"paths"`
	Check struct {
		Body       bool `json:"body"`
		StatusCode bool `json:"statusCode"`
		Header     bool `json:"header"`
	} `json:"check"`
}

// Urls はタグと結び付けられたURL情報の構造体。使用回数が多いので分けた
type Urls struct {
	Tag string `json:"tag"`
	URL string `json:"url"`
}
