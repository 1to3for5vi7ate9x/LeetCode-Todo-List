package main

import (
	"fmt"
	"os"
	"os/signal"
	"Todo/internal/reminder"
	"Todo/pkg/db"
	"Todo/pkg/tasks"
	"Todo/pkg/ui"
	"syscall"
)

func main() {
	// Initialize the database
	db.InitializeDatabase()

	// Initialize tasks for the day
	tasks.InitializeDailyTasks(10)

	// Start the interactive UI
	go ui.StartUI()

	// Start the reminder service
	go reminder.StartReminderService()

	// Handle graceful shutdown on force quit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("\nForce quitting the application. See you tomorrow!")
}

