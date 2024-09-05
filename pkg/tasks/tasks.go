package tasks

import (
	"fmt"
	"Todo/pkg/db"
	"sync"
	"time"
)

var tasks []string
var mu sync.Mutex

// InitializeDailyTasks creates a list with empty tasks
func InitializeDailyTasks(num int) {
	mu.Lock()
	defer mu.Unlock()
	tasks = make([]string, num)
	fmt.Println("Daily LeetCode Tasks Initialized. Fill in each slot as you complete them.")
	printTasks()
}

// AddTask adds a task to the specified index and saves to database
func AddTask(index int, description string) {
	mu.Lock()
	defer mu.Unlock()
	if index >= 0 && index < len(tasks) {
		tasks[index] = description
		date := time.Now().Format("2006-01-02")
		db.AddTask(date, index+1, description, "") // Empty solution initially
		fmt.Printf("Task %d updated!\n", index+1)
	} else {
		fmt.Println("Invalid task number. Please choose a number between 1 and", len(tasks))
	}
	printTasks()
}

// GetTasks returns the current list of tasks
func GetTasks() []string {
	mu.Lock()
	defer mu.Unlock()
	return tasks
}

// Print the current state of tasks
func printTasks() {
	fmt.Println("Today's LeetCode Tasks:")
	for i, task := range tasks {
		if task == "" {
			fmt.Printf("%d. [ ]\n", i+1)
		} else {
			fmt.Printf("%d. [X] %s\n", i+1, task)
		}
	}
}

// ViewPastTasks retrieves and displays solved problems from the database
func ViewPastTasks() {
	db.ViewTasks()
}

