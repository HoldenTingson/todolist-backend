package task

import (
	"database/sql"
	"todolist/util"
)

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) GetTaskRepo() ([]Task, error) {
	var tasks []Task

	rows, err := r.DB.Query("select * from task")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var task Task

		if err := rows.Scan(&task.Id, &task.Name, &task.Priority, &task.Deadline); err != nil {
			return nil, err
		}

		task.Deadline, err = util.FormatIndonesianDate(task.Deadline)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *repository) GetTaskByIdRepo(id int) (Task, error) {
	var task Task

	row := r.DB.QueryRow("select id, name, priority, deadline from task where id = ?", id)

	if err := row.Scan(&task.Id, &task.Name, &task.Priority, &task.Deadline); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (r *repository) CreateTaskRepo(taskRequest *TaskRequest) error {
	_, err := r.DB.Exec("insert into task (name, priority, deadline) values (?, ?, ?)", taskRequest.Name, taskRequest.Priority, taskRequest.Deadline)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateTaskRepo(id int, taskRequest *TaskRequest) error {
	_, err := r.DB.Exec("update task set name = ?, priority = ?, deadline = ? where id = ?", taskRequest.Name, taskRequest.Priority, taskRequest.Deadline, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteTaskRepo(id int) error {
	_, err := r.DB.Exec("delete from task where id = ?", id)

	if err != nil {
		return err
	}
	return nil
}
