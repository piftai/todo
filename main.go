package main

import (
	"fmt"
	"os"
	"todo/tasks"
	"strconv"
)

func main() {
	// need to add a for cycle that will run for each word in user-input
	
	
	command := os.Args[1]
	tm, err := tasks.NewTaskManager("tasks.json")
	if err != nil {
		fmt.Println("Error initializing task manager: ", err)
		return
	}

	switch command {
	case "add":
		description := os.Args[2]
		err := tm.AddTask(description)
		if err != nil {
			fmt.Println("Error adding task: ", err)
		}
	case "--help":
		fmt.Println("Hi there! This is CLI-todo\nI will help how to use it: \nthe 1 argument except ./cli-todo is [command] - add, delete, update, list")
	case "list":
		mass, err := tm.ListTasks()
		fmt.Printf("Here is list of your tasks:")
		os.Stdout.Write(mass)
		if err != nil {
			fmt.Println("Error listing tasks: ", err)
		}
	case "update":
		taskIDStr := os.Args[2]
    	taskID, err := strconv.Atoi(taskIDStr)
		_ = err
		fmt.Println(tm.ShowTask(taskID).Description)
	}

}
