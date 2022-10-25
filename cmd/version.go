/*
Copyright © 2022 Daniils Petrovs <thedanpetrov@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/DaniruKun/hoshinovactl/hoshinova"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the Hoshinova server version",
	Run: func(cmd *cobra.Command, args []string) {
		host, err := cmd.Flags().GetString("host")
		cobra.CheckErr(err)
		port, err := cmd.Flags().GetString("port")
		cobra.CheckErr(err)
		client, err := hoshinova.NewClient(host, port)
		cobra.CheckErr(err)

		version, err := client.GetVersion()
		cobra.CheckErr(err)

		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
