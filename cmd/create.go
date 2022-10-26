/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/DaniruKun/hoshinovactl/hoshinova"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new scheduled task on the server",
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		videoUrl, _ := cmd.Flags().GetString("video-url")
		outputDir, _ := cmd.Flags().GetString("output-dir")

		client, _ := hoshinova.NewClient(host, port)

		err := client.CreateTask(videoUrl, outputDir)
		cobra.CheckErr(err)
	},
}

func init() {
	taskCmd.AddCommand(createCmd)

	createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().String("video-url", "", "The video URL you want to record")
	createCmd.Flags().String("output-dir", "./videos/moona", "The directory where to save recording")
	createCmd.MarkFlagRequired("video-url")
}
