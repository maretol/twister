package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version  string
	Revision string
)

func init() {
	RootCmd.AddCommand(newVersionCmd())
}

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("versin: 0.0.0, revision: -\n")
		},
	}

	return cmd
}
