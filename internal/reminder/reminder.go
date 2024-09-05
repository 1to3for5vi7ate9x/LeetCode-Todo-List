package reminder

import (
	"fmt"
	"time"
)

// StartReminderService triggers notifications periodically
func StartReminderService() {
	for {
		time.Sleep(2 * time.Hour) // Remind every 2 hours
		fmt.Println("🔔 Reminder: Complete your 10 LeetCode problems for today! Keep going!")
	}
}

