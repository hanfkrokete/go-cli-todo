package main

import (
	"fmt"
	"os"
	"strconv"

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
			ID:    task.NextID(tasks),
			Title: title,
			Done:  false,
		}

		tasks = append(tasks, newTask)

		storage.Save(filePath, tasks)

		fmt.Println("✅ Task saved:", title)
	case "done":

		if len(os.Args) < 3 {
			fmt.Println("Usage: todo done <id>")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		tasks, _ := storage.Load(filePath)

		found := false

		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Done = true
				found = true
				fmt.Println("✅ Task marked done:", tasks[i].Title)
				break
			}
		}
		if !found {
			fmt.Println("Task not found")
			return
		}

		storage.Save(filePath, tasks)

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

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo delete <id>")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		tasks, _ := storage.Load(filePath)
		newTasks := make([]task.Task, 0)
		found := false

		for _, t := range tasks {
			if t.ID == id {
				found = true
				continue
			}
			newTasks = append(newTasks, t)
		}

		if !found {
			fmt.Println("Task not found")
			return
		}

		storage.Save(filePath, newTasks)

		fmt.Println("Task deleted:", id)

	default:
		fmt.Println("Unknown command:", command)
	}
}
