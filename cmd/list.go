/*
Copyright © 2025 Svirin Artyom <emotionlesscode@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print the list of all the tasks.",
	Long:  `Print the list of all the tasks.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := LoadTasks()
		if err != nil {
			return err
		}

		if len(tasks.Tasks) == 0 {
			fmt.Println("A list of tasks is empty.")
			return nil
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetStyle(table.StyleColoredBlueWhiteOnBlack)
		t.AppendHeader(table.Row{"#", "Task", "Completed", "Created At", "Completed At"})
		for i, task := range tasks.Tasks {
			t.AppendRow([]interface{}{i + 1, task.Title, task.Completed, task.CreatedAt, task.CompletedAt})
		}
		t.Render()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
