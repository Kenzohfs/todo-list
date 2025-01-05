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

func (r *TaskRepo) Get(page int, perPage int, sortBy string, sortDirection string, filter string, status string) *[]models.Task {
	var tasks []models.Task

	query := r.db.Order(`"` + sortBy + `"` + " " + sortDirection).Limit(perPage).Offset(page * perPage)

	if filter != "" {
		query.Where(`"desc" LIKE ?`, "%"+filter+"%")
	}

	if status != "" {
		query.Where(`status = ?`, status)
	}

	query.Find(&tasks)

	return &tasks
}
