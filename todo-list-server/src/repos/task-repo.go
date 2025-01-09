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

func (r *TaskRepo) Get(workspaceId int, page int, perPage int, sortBy string, sortDirection string, filter string, status string) *[]models.Task {
	var tasks []models.Task

	query := r.db.Order(`"`+sortBy+`"`+" "+sortDirection).Limit(perPage).Offset(page*perPage).Where("workspace_id = ?", workspaceId)

	if filter != "" {
		query.Where(`"desc" LIKE ?`, "%"+filter+"%")
	}

	if status != "" {
		query.Where(`status = ?`, status)
	}

	query.Preload("Workspace").Find(&tasks)

	return &tasks
}

func (r *TaskRepo) GetById(id uint) *models.Task {
	var task models.Task

	r.db.Preload("Workspace").First(&task, id)

	return &task
}

func (r *TaskRepo) Create(task *models.Task) {
	r.db.Save(&task)
}

func (r *TaskRepo) Edit(task *models.Task) {
	r.db.Save(&task)
}

func (r *TaskRepo) Delete(id uint) {
	r.db.Delete(&models.Task{}, id)
}
