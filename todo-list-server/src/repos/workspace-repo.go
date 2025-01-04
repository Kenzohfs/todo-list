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

func (r *WorkspaceRepo) Get(page int, perPage int, sortBy string, sortDirection string, filter string) *[]models.Workspace {
	var workspaces []models.Workspace

	r.db.Find(&workspaces).Where("name LIKE ?", filter).Order(sortBy + " " + sortDirection).Limit(perPage).Offset(page * perPage)

	return &workspaces
}
