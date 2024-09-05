# LeetCode Todo List

This project is a Go-based command-line tool designed to help you stay on track with your daily goal of solving LeetCode problems. It serves as an interactive and persistent to-do list that annoyingly reminds you to complete at least 10 LeetCode problems every day for 21 days. The application is intended to run continuously on your Mac, with the only way to close it being a force quit.

## Features

- **Daily Task List:** Automatically initializes with 10 blank tasks each day.
- **Persistent Tracking:** Tracks your progress over time and saves your completed tasks in a database.
- **Interactive UI:** Provides a simple command-line interface to add tasks, view your progress, and see past completions.
- **Annoying Reminders:** Runs persistently and continuously reminds you of your daily goals.

## Getting Started

### Prerequisites

- **Go**: Ensure that Go is installed on your system. You can download it from [Go's official website](https://golang.org/dl/).
- **SQLite**: The project uses SQLite for persistent storage.

### Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/1to3for5vi7ate9x/Todo-everyday.git
   cd Todo-everyday
   ```

2. **Initialize and Update Dependencies:**

   Initialize the Go module and install any dependencies:

   ```bash
   go mod tidy
   ```

3. **Build the Application:**

   Build the application executable:

   ```bash
   go build -o todo_app ./cmd/Todo/main.go
   ```

### Running the Application

To start the application, run the built executable:

```bash
./todo_app
```

### Using the Application

Once running, the application will display a command prompt where you can interact with your daily tasks. Available commands:

- **Add a Task:** `add <task number> <description>`
- **View Past Tasks:** `view`
- **Quit Forcefully:** Close the terminal window or force quit.

Example:

```text
Commands: add <task number> <description> | view (to view past tasks) | q (to quit forcefully)
add 1 "Two Sum: Solve using HashMap"
```

### Database Management

The application uses SQLite to keep track of all tasks you've completed. The database file is stored in the project directory. To view all past completed tasks, use the `view` command within the application.

## Project Structure

```
Todo-everyday/
├── cmd/
│   └── Todo/
│       └── main.go        # Main entry point of the application
├── pkg/
│   ├── config/            # Configuration management
│   ├── tasks/             # Task management and logic
│   ├── ui/                # User interface and command handling
│   └── db/                # Database interactions
├── internal/
│   └── reminder/          # Reminder service logic
├── go.mod                 # Go module file
├── README.md              # Project README file
└── .gitignore             # Git ignore file
```

## Contributing

Feel free to fork this project, submit issues, and send pull requests. Your contributions are highly appreciated!

## License

This project is open-source and available under the MIT License.

## Contact

For any questions or suggestions, please reach out to ankitc@krates.ai.

---

Happy coding and stay on top of your LeetCode goals!

