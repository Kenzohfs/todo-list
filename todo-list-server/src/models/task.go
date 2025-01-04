package models

import "gorm.io/gorm"

type TaskStatus string

const (
	NotStarted TaskStatus = "NotStarted"
	Completed  TaskStatus = "Completed"
)

type Task struct {
	gorm.Model
	Desc        string     `json:"desc"`
	Status      TaskStatus `json:"status"`
	WorkspaceID int
	Workspace   Workspace
}
