package cmd

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"todo/internal/data"
)

const dataFile = "tasks.json"

func LoadTasks() (data.TaskList, error) {
	var tasks data.TaskList

	file, err := os.Open(dataFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return tasks, nil
		}
		return tasks, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return tasks, err
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func SaveTasks(tasks data.TaskList) error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	js, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	_, err = file.Write(js)
	if err != nil {
		return err
	}
	return nil
}
