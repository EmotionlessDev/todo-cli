/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del [ID]",
	Short: "Delete a task by ID",
	Long: `Delete a task by ID. 
	For example:
	todo del 1
	Will delete the task with ID 1.
	`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		tasks, err := LoadTasks()
		if err != nil {
			return err
		}
		if id > len(tasks.Tasks) || id <= 0 {
			return ErrInvalidID
		}
		tasks.Tasks = append(tasks.Tasks[:id-1], tasks.Tasks[id:]...)

		err = SaveTasks(tasks)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
