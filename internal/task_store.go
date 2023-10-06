package internal

import "sync"

const (
	StatusToDo       = "В ожидании"
	StatusInProgress = "В процессе"
	StatusDone       = "Выполнено"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
}

type TaskStore struct {
	tasks  map[int]Task
	nextID int
	mu     sync.Mutex
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks:  make(map[int]Task),
		nextID: 1,
	}
}

func (ts *TaskStore) CreateTask(task Task) int {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	task.ID = ts.nextID
	ts.tasks[task.ID] = task
	ts.nextID++

	return task.ID
}

func (ts *TaskStore) ReadTask(id int) (Task, bool) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	task, ok := ts.tasks[id]

	return task, ok
}
