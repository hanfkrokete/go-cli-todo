package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ")
		fmt.Println("  todo add \"task\"")
		return
	}

	command := os.Args[1]

	if command == "add" {

		if len(os.Args) < 3 {
			fmt.Println("Error: missing task tittle")
			return
		}

		tittle := os.Args[2]

		fmt.Println("✅ Added task:", tittle)
		return
	}

	fmt.Println("Unknown command:", command)

}
