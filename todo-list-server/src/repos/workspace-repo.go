package repos

import (
	"github.com/kenzohfs/m/src/models"
	"gorm.io/gorm"
)

type WorkspaceRepo struct {
	db *gorm.DB
}

func NewWorkspaceRepo(db *gorm.DB) *WorkspaceRepo {
	return &WorkspaceRepo{db: db}
}

func (r *WorkspaceRepo) Create(workspace *models.Workspace) {
	r.db.Save(&workspace)
}
