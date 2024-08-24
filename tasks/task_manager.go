package tasks

import (
	"encoding/json"
	"os"
	"time"
)

type TaskManager struct {
	tasks    []Task
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

// this func should be written at end of every func that communicate with tm object.
// func saving changing
func (tm *TaskManager) saveTasks() error {
	data, err := json.MarshalIndent(tm.tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(tm.filePath, data, 0644)
}

// added a task into file
func (tm *TaskManager) AddTask(description string) error {
	task := Task{
		ID:          len(tm.tasks) + 1, // there a existing a bug when you can create two tasks with same id
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tm.tasks = append(tm.tasks, task)
	return tm.saveTasks()
}

// return a file tasks.json
func (tm *TaskManager) ListTasks() ([]byte, error) {
	return os.ReadFile(tm.filePath)
}

// return updated task by id and new description
func (tm *TaskManager) UpdateTask(taskID int, description string) *Task {
	for i := range tm.tasks {
		if tm.tasks[i].ID == taskID {
			tm.tasks[i].Description = description
			tm.tasks[i].UpdatedAt = time.Now()
			anyError := tm.saveTasks()
			if anyError != nil {
				return nil
			}
			return &tm.tasks[i]
		}
	}
	return nil
}

// return pointer of task by id
func (tm *TaskManager) ShowTask(taskID int) *Task {
	for _, task := range tm.tasks {
		if task.ID == taskID {
			return &task
		}
	}
	return nil
}

func (tm *TaskManager) DeleteTask(taskID int) bool {
	for i := 0; i < len(tm.tasks); i++ {
		if tm.tasks[i].ID == taskID {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			tm.saveTasks()
			return true
		}
	}
	return false
}
