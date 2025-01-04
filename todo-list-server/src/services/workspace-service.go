package services

import (
	"github.com/kenzohfs/m/src/models"
	"github.com/kenzohfs/m/src/repos"
)

type WorkspaceService struct {
	workspaceRepo *repos.WorkspaceRepo
}

func NewWorkspaceService(workspaceRepo *repos.WorkspaceRepo) *WorkspaceService {
	return &WorkspaceService{workspaceRepo: workspaceRepo}
}

func (s *WorkspaceService) Create(workspace *models.Workspace) {
	s.workspaceRepo.Create(workspace)
}

func (r *WorkspaceService) Get(page int, perPage int, sortBy string, sortDirection string, filter string) *[]models.Workspace {
	return r.workspaceRepo.Get(page, perPage, sortBy, sortDirection, filter)
}

func (r *WorkspaceService) GetById(id uint) *models.Workspace {
	return r.workspaceRepo.GetById(id)
}
