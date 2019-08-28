package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(configCmd())
}

func configCmd() *cobra.Command {
	baseCmd := &cobra.Command{
		Use:   "config [command]",
		Short: "config command",
		Long:  "You can create config template and will be able to edit config file",
		// Run:   hoge,
	}
	baseCmd.AddCommand(configCreateCmd())
	baseCmd.AddCommand(configEditCmd())
	baseCmd.AddCommand(configShowCmd())
	return baseCmd
}

func configCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create config template file",
		Long:  "Create new config file in current directory",
		Run:   configCreating,
	}
	return cmd
}

func configEditCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit [config_key] [config_value]",
		Short: "edit config value",
		Long:  "まだ実装してない。ごめん",
		Run:   empty,
	}
	return cmd
}

func configShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "showing config",
		Run:   configShowing,
	}
	return cmd
}

func configCreating(cmd *cobra.Command, args []string) {
	if IsConfigExist {
		fmt.Print("Config file is exist. Do you allow to overwrite new config? [y/N] : ")
		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			log.Fatal(err)
		}
		if response != "y" {
			return
		}
	}
	file, err := os.OpenFile("twisterconf.json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	Config = ConfigBase{}
	Config.Urls = []string{""}
	Config.Paths = []string{""}
	Config.Check.Body = true
	Config.Check.Header = true
	Config.Check.StatusCode = true
	jsonBytes, err := json.Marshal(Config)
	var buf bytes.Buffer
	json.Indent(&buf, jsonBytes, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
}

func configShowing(cmd *cobra.Command, args []string) {
	if !IsConfigExist {
		fmt.Println("Config file not found!")
		return
	}
	urls := Config.Urls
	for num := range urls {
		fmt.Printf("URL%d : %s\n", num, urls[num])
	}
	paths := Config.Paths
	for num := range paths {
		fmt.Printf("Path : %s\n", paths[num])
	}
	fmt.Println("Checking")
	fmt.Printf("Header     : %t\n", Config.Check.Header)
	fmt.Printf("Body       : %t\n", Config.Check.Body)
	fmt.Printf("StatusCode : %t\n", Config.Check.StatusCode)
}

func empty(cmd *cobra.Command, args []string) {
	fmt.Printf("未実装\n")
}
