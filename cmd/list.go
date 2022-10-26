/*
Copyright © 2022 Daniils Petrovs <thedanpetrov@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/DaniruKun/hoshinovactl/hoshinova"
	"github.com/InVisionApp/tabular"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List scheduled recording tasks",
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		client, _ := hoshinova.NewClient(host, port)
		taskItems, err := client.GetTasks()
		cobra.CheckErr(err)

		tab := tabular.New()

		tab.Col("videoid", "VideoID", 12)
		tab.Col("state", "State", 14)
		tab.Col("size", "Total Size", 16)

		format := tab.Print("*")

		// header := width.Widen.String(fmt.Sprintf("%-40s %-10s %-12s\n", "Title", "State", "Progress"))
		// headerSep := width.Widen.String(fmt.Sprintf(strings.Repeat("-", titleLen) + " " + strings.Repeat("-", stateLen) + " " + strings.Repeat("-", progressLen)))

		lo.ForEach(taskItems, func(taskItem hoshinova.TaskItem, _ int) {
			fmt.Printf(format, taskItem.VideoId, hoshinova.EscapeStatus(taskItem.Status.State), taskItem.TotalSize)
		})
	},
}

func init() {
	taskCmd.AddCommand(listCmd)
}
