package model

import (
	"errors"
	"fmt"
)

type Task struct {
	name string
	description string
}

func (t Task) Str() string {
	return fmt.Sprintf("\n\t name: %s \n\t description: %s  \n } \n ", t.name, t.description)
}

var tasks = make([]Task, 0, 1000)

func CreateTask(name, desc string) (int, error) {
	if(len(name) == 0) {
		return -1, errors.New("The name should be provided")
	}
	tasks = append(tasks, Task{name, desc})

	return len(tasks) - 1, nil
}

func UpdateTask(id int, name, desc string) (int, error) {
	if(id < 0) {
		return -1, errors.New("Id is invalid")
	} else if(len(name) == 0) {
		return -1, errors.New("The name should be provided")
	} else {
		tasks[id].name = name
		tasks[id].description = desc
		return id, nil
	}
}


func GetAllTasks() (string, error) {
	str := "";

	for i, v := range tasks {
		str += fmt.Sprintf("{ \n\t id: %d %s", i, v.Str())
	}

	return str, nil
}







