/*
Copyright © 2024 Svirin Artyom <emotionlesscode@gmail.com>
*/
package cmd

import (
	"os"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Basic todo list",
	Long:  `This is a basic todo list application.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
