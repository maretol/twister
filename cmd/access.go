package cmd

import (
	"github.com/spf13/cobra"

	"github.com/maretol/twister/logic"
)

func init() {
	RootCmd.AddCommand(accessCmd())
}

var (
	header     bool
	body       bool
	statusCode bool
	output     string
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
	accessAllURL(urls, paths)
}

func accessAllURL(urls []string, pathes []string) {
	for _, path := range pathes {
		booster := logic.Booster{}
		booster.SetFullURL(path, urls)
		booster.AllAccess()
		logic.Output(booster, Config.Check.StatusCode, Config.Check.Header, Config.Check.Body)
	}
}
