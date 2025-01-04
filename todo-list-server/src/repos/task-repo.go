package repos

import (
	"github.com/kenzohfs/m/src/models"
	"gorm.io/gorm"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) Create(task *models.Task) {
	r.db.Save(&task)
}
