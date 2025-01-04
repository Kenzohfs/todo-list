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
