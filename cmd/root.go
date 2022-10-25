/*
Copyright © 2022 Daniils Petrovs <thedanpetrov@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hoshinovactl",
	Short: "CLI to interact with the Hoshinova stream recording server",
	Long:  `hoshinovactl allows you to schedule recording tasks, list currently scheduled tasks, and more.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Set defaults
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", "1104")

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hoshinovactl.yaml)")
	rootCmd.PersistentFlags().String("host", viper.GetString("host"), "hostname or address of Hoshinova server")
	rootCmd.PersistentFlags().StringP("port", "p", viper.GetString("port"), "port on which Hoshinova API is accessible")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".hoshinovactl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".hoshinovactl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

}
