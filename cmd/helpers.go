package cmd

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"todo/internal/data"
)

func GetDataFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Library", "Application Support", "todo", "data.json"), nil
}

func LoadTasks() (data.TaskList, error) {
	var tasks data.TaskList

	dataFile, err := GetDataFilePath()
	if err != nil {
		return tasks, err
	}

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
	dataFile, err := GetDataFilePath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(dataFile)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

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
