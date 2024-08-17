package main

import (
	"fmt"
	"os"
	"todo/tasks"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli-todo [command] [description]")
	}

	command := os.Args[1]
	tm, err := tasks.NewTaskManager("tasks.json")
	if err != nil {
		fmt.Println("Error initializing task manager: ", err)
		return
	}

	switch command {
	case "add":
		fmt.Println("you are here")
		description := os.Args[2]
		err := tm.AddTask(description)
		if err != nil {
			fmt.Println("Error adding task: ", err)
		} else {
			fmt.Println("all is ok")
		}
	case "--help":
		fmt.Println("Hi there! This is CLI-todo\nI will help how to use it: \nthe 1 argument except ./cli-todo is [command] - add, delete, update, list")
	case "list":
		mass, err := tm.ListTasks()
		fmt.Println("Here is list of your tasks:\n")
		os.Stdout.Write(mass)
		if err != nil {
			fmt.Println("Error listing tasks: ", err)
		}
	}
}
// i like to code
