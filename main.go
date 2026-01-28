package main

import (
	"fmt"
	"os"

	"github.com/hanfkrokete/go-cli-todo/internal/storage"
	"github.com/hanfkrokete/go-cli-todo/internal/task"
)

const filePath = "data/tasks.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ")
		fmt.Println("  todo add \"task\"")
		fmt.Println("  todo list")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: missing task title")
			return
		}

		title := os.Args[2]

		tasks, _ := storage.Load(filePath)

		newTask := task.Task{
			ID:    len(tasks) + 1,
			Title: title,
			Done:  false,
		}

		tasks = append(tasks, newTask)

		storage.Save(filePath, tasks)

		fmt.Println("✅ Task saved:", title)

	case "list":
		tasks, _ := storage.Load(filePath)

		if len(tasks) == 0 {
			fmt.Println("No tasks yet ✅")
			return
		}

		for _, t := range tasks {
			status := "[ ]"
			if t.Done {
				status = "[x]"
			}
			fmt.Printf("  %s %d: %s\n", status, t.ID, t.Title)
		}
	default:
		fmt.Println("Unknown command:", command)
	}
}
