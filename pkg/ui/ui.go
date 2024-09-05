package ui

import (
	"bufio"
	"fmt"
	"os"
	"Todo/pkg/tasks"
	"strconv"
	"strings"
)

// StartUI begins the interactive terminal UI
func StartUI() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nCommands: add <task number> <description> | view (to view past tasks) | q (to quit forcefully)")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "view" {
			tasks.ViewPastTasks()
			continue
		}
		parts := strings.SplitN(input, " ", 3)

		if len(parts) < 3 || parts[0] != "add" {
			fmt.Println("Invalid command. Use 'add <task number> <description>' or 'view' or 'q'")
			continue
		}

		index, err := strconv.Atoi(parts[1])
		if err != nil || index < 1 || index > len(tasks.GetTasks()) {
			fmt.Println("Invalid task number. Please enter a number between 1 and", len(tasks.GetTasks()))
			continue
		}

		tasks.AddTask(index-1, parts[2])
	}
}

