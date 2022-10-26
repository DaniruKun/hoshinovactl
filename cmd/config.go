/*
Copyright © 2022 Daniils Petrovs <thedanpetrov@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "view and manage server configuration",
}

func init() {
	rootCmd.AddCommand(configCmd)
}
