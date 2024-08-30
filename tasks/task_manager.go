package tasks

type TaskManager struct {
	tasks    []Task
	filePath string
}

func NewTaskManager(filepath string) (*TaskManager, error) {
	tm := &TaskManager{filePath: filepath}

}

func (tm *TaskManager) loadTasks() error {

}
