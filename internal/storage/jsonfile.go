package storage

import (
	"encoding/json"
	"os"

	"github.com/hanfkrokete/go-cli-todo/internal/task"
)

func Load(path string) ([]task.Task, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return []task.Task{}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tasks []task.Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil

}

func Save(path string, tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
