package service

import (
	"task-rpc-server/model"
)

type Args struct {
	id int
	name, description string
}

type TaskService struct{}

func (t *TaskService) CreateTask(args *Args, reply *int) error {
	id, err := model.CreateTask(args.name, args.description)

	if(err!= nil) {
		return err
	}

	*reply = id

	return nil
}

func (t *TaskService) GetAllTasks(args *Args, reply *string) error {

	tasks, err := model.GetAllTasks()

	if(err != nil) {
		return err
	}

	*reply = tasks

	return nil
}