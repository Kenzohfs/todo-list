package models

import "gorm.io/gorm"

type Workspace struct {
	gorm.Model
	Name string `json:"name"`
}
