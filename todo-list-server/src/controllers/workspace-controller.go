package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenzohfs/m/src/models"
	"github.com/kenzohfs/m/src/services"
)

type WorkspaceController struct {
	workspaceService *services.WorkspaceService
}

func NewWorkspaceController(workspaceService *services.WorkspaceService) *WorkspaceController {
	return &WorkspaceController{workspaceService: workspaceService}
}

func (c *WorkspaceController) Create(ctx *gin.Context) {
	var workspace models.Workspace

	if err := ctx.ShouldBindBodyWithJSON(&workspace); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.workspaceService.Create(&workspace)
	ctx.JSON(http.StatusCreated, workspace)
}
