package main

import (
	"errors"
)

type Task struct {
	ID   int
	Name string
}

var tasks []Task
var nextID int

func AddTask(name string) Task {
	nextID++
	task := Task{ID: nextID, Name: name}
	tasks = append(tasks, task)
	return task
}

func DeleteTask(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

func ListTasks() []Task {
	return tasks
}
