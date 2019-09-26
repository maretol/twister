package cmd

import (
	"github.com/spf13/cobra"

	"github.com/maretol/twister/logic"
)

func init() {
	RootCmd.AddCommand(accessCmd())
}

var (
	header       bool
	body         bool
	statusCode   bool
	headerIgnore []string
	output       string
)

func accessCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "access",
		Short: "access config each url and each path",
		Long:  "access the url in config file and each paths.\nflags take priority of config file.",
		Run:   access,
	}
	cmd.Flags().BoolVarP(&header, "header", "e", Config.Check.Header, "When this option use, checking result of each header match")
	cmd.Flags().BoolVarP(&body, "body", "b", Config.Check.Body, "When this option use, checking result of each body match")
	cmd.Flags().BoolVarP(&statusCode, "statuscode", "s", Config.Check.StatusCode, "When this option use, checking result of HTTP status code match")
	cmd.Flags().StringVarP(&output, "output-file", "o", "", "setting result output file(未実装)")
	return cmd
}

func access(cmd *cobra.Command, args []string) {
	urls := Config.Urls
	paths := Config.Paths
	header = Config.Check.Header
	body = Config.Check.Body
	statusCode = Config.Check.StatusCode
	headerIgnore = Config.Header.Ignore
	accessAllURL(urls, paths)
}

func accessAllURL(urls []Urls, pathes []string) {
	for _, path := range pathes {
		booster := logic.Booster{}
		passURL := make([]logic.GetURL, len(urls))
		for num, obj := range urls {
			passURL[num].URL = obj.URL
			passURL[num].Tag = obj.Tag
		}
		booster.SetFullURL(path, passURL)
		booster.AllAccess()
		booster.Checker.StatusCode = statusCode
		booster.Checker.Header = header
		booster.Checker.Body = body
		booster.HeaderIgnore = headerIgnore
		logic.Output(booster)
	}
}
