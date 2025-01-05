package services

import (
	"github.com/kenzohfs/m/src/models"
	"github.com/kenzohfs/m/src/repos"
)

type TaskService struct {
	taskRepo *repos.TaskRepo
}

func NewTaskService(taskRepo *repos.TaskRepo) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) Create(task *models.Task) {
	s.taskRepo.Create(task)
}

func (r *TaskService) Get(page int, perPage int, sortBy string, sortDirection string, filter string, status string) *[]models.Task {
	return r.taskRepo.Get(page, perPage, sortBy, sortDirection, filter, status)
}
