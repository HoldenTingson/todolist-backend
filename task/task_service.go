package task

type service struct {
	Repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		Repository: repository,
	}
}

func (s *service) GetTaskService() ([]Task, error) {
	tasks, err := s.Repository.GetTaskRepo()

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *service) GetTaskByIdService(id int) (Task, error) {
	task, err := s.Repository.GetTaskByIdRepo(id)

	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *service) CreateTaskService(taskRequest *TaskRequest) error {
	if err := s.Repository.CreateTaskRepo(taskRequest); err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateTaskService(id int, taskRequest *TaskRequest) error {
	if err := s.Repository.UpdateTaskRepo(id, taskRequest); err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteTaskService(id int) error {
	if err := s.Repository.DeleteTaskRepo(id); err != nil {
		return err
	}

	return nil
}
