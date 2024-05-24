package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("1. Add Task")
		fmt.Println("2. Delete Task")
		fmt.Println("3. List Tasks")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task name: ")
			scanner.Scan()
			name := scanner.Text()
			task := AddTask(name)
			fmt.Printf("Added task: %d - %s\n", task.ID, task.Name)
		case "2":
			fmt.Print("Enter task ID to delete: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid task ID")
				continue
			}
			err = DeleteTask(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task deleted")
			}
		case "3":
			tasks := ListTasks()
			if len(tasks) == 0 {
				fmt.Println("No tasks available")
			} else {
				fmt.Println("Tasks:")
				for _, task := range tasks {
					fmt.Printf("%d: %s\n", task.ID, task.Name)
				}
			}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
