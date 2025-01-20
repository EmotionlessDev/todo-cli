/*
Copyright © 2025 Svirin Artyom <emotionlesscode@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"todo/internal/data"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [taskName]",
	Short: "add a new task",
	Long: `add a new task to the list of tasks
	if taskName is not provided, it will be read from stdin
	`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var title string

		if len(args) > 0 {
			title = args[0]
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			if err := scanner.Err(); err != nil {
				return err
			}

			if len(scanner.Text()) == 0 {
				return fmt.Errorf("task name can't be empty")
			}

			title = scanner.Text()
		}
		title = strings.TrimSpace(title)

		tasks, err := LoadTasks()
		if err != nil {
			return err
		}

		newTask := data.Task{
			Title:       title,
			Completed:   false,
			CreatedAt:   time.Now().Format(time.RFC3339),
			CompletedAt: "",
		}

		tasks.Tasks = append(tasks.Tasks, newTask)

		err = SaveTasks(tasks)
		if err != nil {
			return err
		}

		fmt.Printf("Task '%s' added\n", title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
