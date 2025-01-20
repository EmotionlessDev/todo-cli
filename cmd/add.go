/*
Copyright © 2025 Svirin Artyom <emotionlesscode@gmail.com>
*/
package cmd

import (
	"fmt"
	"todo/internal/data"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [taskName]",
	Short: "add a new task",
	Long:  `add a new task to the list of tasks`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		title := args[0]

		tasks, err := LoadTasks()
		if err != nil {
			return err
		}

		newTask := data.Task{
			ID:        len(tasks.Tasks) + 1,
			Title:     title,
			Completed: false,
		}

		tasks.Tasks = append(tasks.Tasks, newTask)

		err = SaveTasks(tasks)
		if err != nil {
			return err
		}

		fmt.Printf("Задача '%s' добавлена!\n", title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
