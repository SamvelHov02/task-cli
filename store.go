package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Store struct {
	FilePath  string
	IdCounter uint32
}

type Database struct {
	IdCounter uint32 `json:"counter"`
	Tasks     []Task `json:"tasks"`
}

type Task struct {
	Id          uint32 `json:"id"`
	Description string `json:"desc"`
	Status      Status `json:"status"`
}

func GetHomeDir() (string, error) {
	base, _ := os.UserConfigDir()
	dir := filepath.Join(base, "task-cli")

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}

	return dir, nil
}

func ReadTasks(s *Store) (*Database, error) {
	jsonFile, err := os.ReadFile(s.FilePath)

	var database Database

	if errors.Is(err, os.ErrNotExist) {
		// Json file doesn't exist and needs to be created
		database.IdCounter = 0
		database.Tasks = []Task{}

		return &database, nil
	}

	err = json.Unmarshal(jsonFile, &database)

	s.IdCounter = database.IdCounter

	if err != nil {
		return nil, err
	}

	return &database, nil
}

func WriteDatabase(d Database, s *Store) error {
	byteValue, err := json.Marshal(d)

	if err != nil {
		return err
	}

	os.WriteFile(s.FilePath, byteValue, 0o644)

	return nil
}

func (s *Store) InitilizeStore() error {
	dir, err := GetHomeDir()

	if err != nil {
		return err
	}

	s.FilePath = filepath.Join(dir, "tasks.json")
	return nil
}
