package main

import (
	"slices"
	"strconv"
)

type Status string

const (
	StatusDone       Status = "done"
	StatusInProgress Status = "in-progress"
	StatusTodo       Status = "todo"
)

type Command interface {
	Name() string
	Run(arguments []string, s *Store) ([]Task, error)
}

type AddCommand struct {
}

type UpdateCommand struct {
}

type DeleteCommand struct {
}

type MarkCommand struct {
}

type ListCommand struct {
	FilterBy []Status
}

func (c AddCommand) Name() string {
	return "Adds a task to the list, usage : task-cli add 'Finish the project'"
}

func (c AddCommand) Run(arguments []string, s *Store) ([]Task, error) {
	database, err := ReadTasks(s)

	if err != nil {
		return nil, err
	}
	NewCount := database.IdCounter + 1
	database.IdCounter = NewCount

	task := Task{Id: NewCount, Status: StatusTodo, Description: arguments[0]}
	database.Tasks = append(database.Tasks, task)

	err = WriteDatabase(*database, s)

	return nil, err
}

func (c ListCommand) Name() string {
	return "Lists all the tasks. Can filter by their status, default shows the active and todo tasks"
}

func (c ListCommand) Run(arguments []string, s *Store) ([]Task, error) {
	// Default behaviour won't need an argument
	if len(arguments) > 0 {
		c.FilterBy = []Status{Status(arguments[0])}
	} else {
		c.FilterBy = []Status{StatusInProgress, StatusTodo}
	}
	database, err := ReadTasks(s)

	if err != nil {
		return nil, err
	}

	InterestTasks := make([]Task, 0, len(database.Tasks))

	for _, t := range database.Tasks {
		if slices.Contains(c.FilterBy, t.Status) {
			InterestTasks = append(InterestTasks, t)
		}
	}

	return InterestTasks, nil
}

func (c UpdateCommand) Name() string {
	return "Updates a task, usage 'task-cli update <id> 'update the website''"
}

func (c UpdateCommand) Run(arguments []string, s *Store) ([]Task, error) {
	id, err := strconv.Atoi(arguments[0])
	if err != nil {
		return nil, err
	}

	id2 := uint32(id)

	database, err := ReadTasks(s)

	if err != nil {
		return nil, err
	}

	for i, t := range database.Tasks {
		if t.Id == id2 {
			database.Tasks[i].Description = arguments[1]
		}
	}

	err = WriteDatabase(*database, s)

	return nil, err
}

func (c DeleteCommand) Name() string {
	return "Deletes a task from database, usage 'task-cli delete <id>'"
}

func (c DeleteCommand) Run(arguments []string, s *Store) ([]Task, error) {
	id, err := strconv.Atoi(arguments[0])
	if err != nil {
		return nil, err
	}

	id2 := uint32(id)

	database, err := ReadTasks(s)

	if err != nil {
		return nil, err
	}

	toRemoveidx := -1

	for i, t := range database.Tasks {
		if t.Id == id2 {
			toRemoveidx = i
		}
	}

	if toRemoveidx > -1 {
		database.Tasks = append(database.Tasks[:toRemoveidx], database.Tasks[toRemoveidx+1:]...)
		err = WriteDatabase(*database, s)

		return nil, err
	}

	return nil, err
}

func (c MarkCommand) Name() string {
	return "Changes status of a task, usage 'task-cli mark <id> done'"
}

func (c MarkCommand) Run(arguments []string, s *Store) ([]Task, error) {
	// Check that there is correct number of arguments
	id, err := strconv.Atoi(arguments[0])

	id2 := uint32(id)

	if err != nil {
		return nil, err
	}

	NewStatus := Status(arguments[1])

	if NewStatus != StatusDone && NewStatus != StatusInProgress && NewStatus != StatusTodo {
		return nil, nil
	}

	database, err := ReadTasks(s)
	if err != nil {
		return nil, err
	}

	for i, t := range database.Tasks {
		if t.Id == id2 {
			database.Tasks[i].Status = NewStatus
		}
	}

	err = WriteDatabase(*database, s)

	return nil, err
}
