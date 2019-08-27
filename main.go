package main

import (
	"fmt"
	"os"

	"github.com/maretol/twister/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}
