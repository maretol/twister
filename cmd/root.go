package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var rootCmd = &cobra.Command{
	Use:   "hoge",
	Short: "fuga",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hoge opening")
	},
}
