package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitializeDatabase sets up the SQLite database
func InitializeDatabase() {
	// Create data directory if not exists
	dataDir := filepath.Join(os.Getenv("HOME"), ".todo_project")
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create data directory: %v", err)
	}

	// Open database connection
	dbPath := filepath.Join(dataDir, "tasks.db")
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Create tasks table if not exists
	createTable := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TEXT,
		task_number INTEGER,
		description TEXT,
		solution TEXT
	);`

	if _, err := database.Exec(createTable); err != nil {
		log.Fatalf("Failed to create tasks table: %v", err)
	}

	DB = database
	fmt.Println("Database initialized successfully.")
}

// AddTask inserts a new task into the database
func AddTask(date string, taskNumber int, description string, solution string) {
	statement, err := DB.Prepare("INSERT INTO tasks (date, task_number, description, solution) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Printf("Failed to prepare insert statement: %v", err)
		return
	}
	defer statement.Close()

	_, err = statement.Exec(date, taskNumber, description, solution)
	if err != nil {
		log.Printf("Failed to insert task: %v", err)
	} else {
		fmt.Println("Task saved to the database.")
	}
}

// ViewTasks retrieves and prints all tasks from the database
func ViewTasks() {
	rows, err := DB.Query("SELECT id, date, task_number, description, solution FROM tasks")
	if err != nil {
		log.Printf("Failed to query tasks: %v", err)
		return
	}
	defer rows.Close()

	fmt.Println("Solved LeetCode Problems:")
	for rows.Next() {
		var id, taskNumber int
		var date, description, solution string
		if err := rows.Scan(&id, &date, &taskNumber, &description, &solution); err != nil {
			log.Printf("Failed to scan row: %v", err)
			continue
		}
		fmt.Printf("ID: %d, Date: %s, Task #%d: %s\nSolution:\n%s\n", id, date, taskNumber, description, solution)
	}
}

