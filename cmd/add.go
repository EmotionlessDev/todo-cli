/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [taskName]",
	Short: "add a new task",
	Long:  `add a new task to the list of tasks`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		fmt.Println("add called with task:", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
