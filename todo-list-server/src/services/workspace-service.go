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
