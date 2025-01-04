package main

import (
	"testing"

	"github.com/kenzohfs/m/src/databases"
	"github.com/kenzohfs/m/src/migrations"
	"github.com/kenzohfs/m/src/repos"
	"github.com/kenzohfs/m/src/services"
)

func TestGetWorkspaceById(t *testing.T) {
	db := databases.Connect()
	migrations.Migrate(db)

	workspaceRepo := repos.NewWorkspaceRepo(db)
	workspaceService := services.NewWorkspaceService(workspaceRepo)

	workspaceService.GetById(uint(1))
}
