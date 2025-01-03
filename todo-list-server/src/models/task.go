package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Desc        string `json:"desc"`
	WorkspaceID int
	Workspace   Workspace
}
