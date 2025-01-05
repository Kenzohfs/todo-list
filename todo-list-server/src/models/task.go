package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type TaskStatus string

const (
	NotStarted TaskStatus = "NotStarted"
	Completed  TaskStatus = "Completed"
)

type Task struct {
	gorm.Model
	Desc        string     `json:"desc" validate:"nonzero"`
	Status      TaskStatus `json:"status" validate:"nonzero"`
	WorkspaceID int
	Workspace   Workspace
}

func (m *Task) Validate() error {
	if err := validator.Validate(m); err != nil {
		return err
	}

	return nil
}
