package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Text      string
	Completed bool
}

func main() {
	tasks := []Task{}

	for {
		showMenu()
		option := getUserInput("Enter your choice: ")

		switch option {
			case "1":
				showTasks(tasks)
			case "2":
				addTask(&tasks)
			case "3":
				markTaskCompleted(&tasks)
			case "4":
				saveTasksToFile(tasks)
			case "5":
				fmt.Println("Exiting the ToDo application.")
				return
			default:
				fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func showMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Show Tasks")
	fmt.Println("2. Add Task")
	fmt.Println("3. Mark Task as Completed")
	fmt.Println("4. Save Tasks to File")
	fmt.Println("5. Exit")
}

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func showTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	fmt.Println("Tasks:")
	for i, task := range tasks {
		status := " "
		if task.Completed {
			status = "x"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Text)
	}
}

func addTask(tasks *[]Task) {
	taskText := getUserInput("Enter task description: ")
	*tasks = append(*tasks, Task{Text: taskText})
	fmt.Println("Task added.")
}

func markTaskCompleted(tasks *[]Task) {
	showTasks(*tasks)
	taskIndexStr := getUserInput("Enter task number to mark as completed: ")

	taskIndex, err := parseTaskIndex(taskIndexStr)
	if err != nil || taskIndex < 1 || taskIndex > len(*tasks) {
		fmt.Println("Invalid task number. Please try again.")
		return
	}

	(*tasks)[taskIndex-1].Completed = true
	fmt.Println("Task marked as completed.")
}

func parseTaskIndex(input string) (int, error) {
	return strconv.Atoi(input)
}

func saveTasksToFile(tasks []Task) {
	file, err := os.Create("tasks.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "x"
		}
		file.WriteString(fmt.Sprintf("[%s] %s\n", status, task.Text))
	}

	fmt.Println("Tasks saved to file 'tasks.txt'.")
}
