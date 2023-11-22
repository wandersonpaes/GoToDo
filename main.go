package main

import "fmt"

type TaskStruct struct {
	name        string
	description string
}

func main() {
	var comand int
	tasks := []TaskStruct{}

	for {
		printOption()
		fmt.Scan(&comand)

		if comand == 1 {
			tasks = addNewTask(tasks)
		} else if comand == 2 {
			fmt.Println("2")
		} else if comand == 3 {
			fmt.Println("3")
		} else if comand == 4 {
			break
		}
	}
}

func showTasks(tasks []TaskStruct) {
	for i, task := range tasks {
		fmt.Println("Task ", i, ": ", task.name)
		fmt.Println("Description: ", task.description)
		fmt.Println()
	}
}

func addNewTask(tasks []TaskStruct) []TaskStruct {
	var name, description string

	fmt.Println("Add a name for the task:")
	fmt.Scan(&name)
	fmt.Println("Add a description for the task:")
	fmt.Scan(&description)

	tasks = append(tasks, TaskStruct{
		name:        name,
		description: description,
	})

	return tasks
}

func printOption() {
	fmt.Println("MENU GoToDo")
	fmt.Println("\n1 - Add a new task")
	fmt.Println("2 - Remove a task")
	fmt.Println("3 - Show all tasks")
	fmt.Println("4 - Exit")
}
