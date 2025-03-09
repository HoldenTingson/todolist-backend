package task

type Task struct {
	Id       int
	Name     string
	Priority string
	Deadline string
}

type TaskRequest struct {
	Name     string
	Priority string
	Deadline string
}

type Repository interface {
	GetTaskRepo() ([]Task, error)
	GetTaskByIdRepo(id int) (Task, error)
	CreateTaskRepo(taskRequest *TaskRequest) error
	UpdateTaskRepo(id int, taskRequest *TaskRequest) error
	DeleteTaskRepo(id int) error
}

type Service interface {
	GetTaskService() ([]Task, error)
	GetTaskByIdService(id int) (Task, error)
	CreateTaskService(taskRequest *TaskRequest) error
	UpdateTaskService(id int, taskRequest *TaskRequest) error
	DeleteTaskService(id int) error
}
