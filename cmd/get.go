/*
Copyright © 2022 Daniils Petrovs <thedanpetrov@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/DaniruKun/hoshinovactl/hoshinova"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the server configuration as a JSON string.",
	Run: func(cmd *cobra.Command, args []string) {
		host, err := cmd.Flags().GetString("host")
		cobra.CheckErr(err)
		port, err := cmd.Flags().GetString("port")
		cobra.CheckErr(err)
		client, err := hoshinova.NewClient(host, port)
		cobra.CheckErr(err)

		config, err := client.GetConfig()
		cobra.CheckErr(err)

		fmt.Println(config)
	},
}

func init() {
	configCmd.AddCommand(getCmd)
}
