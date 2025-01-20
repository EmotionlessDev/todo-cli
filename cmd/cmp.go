/*
Copyright © 2025 Svirin Artyom <emotionlesscode@gmail.com>
*/
package cmd

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var cmpCmd = &cobra.Command{
	Use:   "cmp [ID]",
	Short: "Complete the task",
	Long:  `Complete the task by marking it as done.`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var id int

		if len(args) > 0 {
			var err error
			id, err = strconv.Atoi(args[0])
			if err != nil {
				return err
			}
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			if err := scanner.Err(); err != nil {
				return err
			}

			if len(scanner.Text()) == 0 {
				return errors.New("task ID can't be empty")
			}

			var err error
			id, err = strconv.Atoi(scanner.Text())
			if err != nil {
				return err
			}
		}

		tasks, err := LoadTasks()
		if err != nil {
			return err
		}
		if id > len(tasks.Tasks) || id <= 0 {
			return ErrInvalidID
		}
		tasks.Tasks[id-1].Completed = true
		tasks.Tasks[id-1].CompletedAt = time.Now().Format(time.RFC3339)

		err = SaveTasks(tasks)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cmpCmd)
}
