package tasks

import (
	"encoding/json"
	"os"
	"time"
)

type TaskManager struct {
	tasks []Task
	filePath string
}

// constructor of TaskManager structure 
func NewTaskManager(filePath string) (*TaskManager, error) {
	tm := &TaskManager{filePath: filePath}
	err := tm.loadTasks()
	return tm, err
}

// func that creates a new file if its not exist by path of tm.filePath
func (tm *TaskManager) loadTasks() error {
	if _, err := os.Stat(tm.filePath); os.IsNotExist(err) {
		tm.tasks = []Task{}
		return nil
	}
	data, err := os.ReadFile(tm.filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &tm.tasks)
}

// writing a task into file
func (tm *TaskManager) saveTasks() error {
	data, err := json.MarshalIndent(tm.tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(tm.filePath, data, 0644)
}

// added a task into file
func (tm *TaskManager) AddTask(description string) error {
	task := Task {
		ID: len(tm.tasks) + 1,
		Description: description,
		Status: "todo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tm.tasks = append(tm.tasks, task)
	return tm.saveTasks()
}

// return a file tasks.json
func (tm *TaskManager) ListTasks() ([]byte, error) {
	return os.ReadFile(tm.filePath)
}

