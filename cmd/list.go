/*
Copyright © 2025 Svirin Artyom <emotionlesscode@gmail.com>
*/
package cmd

import (
	"fmt"

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

		fmt.Println("List of tasks:")
		for _, task := range tasks.Tasks {
			fmt.Printf("%d. %s\n", task.ID, task.Title)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
