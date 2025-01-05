package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kenzohfs/m/src/models"
	"github.com/kenzohfs/m/src/services"
)

type TaskController struct {
	taskService      *services.TaskService
	workspaceService *services.WorkspaceService
}

func NewTaskController(taskService *services.TaskService, workspaceService *services.WorkspaceService) *TaskController {
	return &TaskController{taskService: taskService, workspaceService: workspaceService}
}

func (c *TaskController) Create(ctx *gin.Context) {
	var task models.Task

	if err := ctx.ShouldBindBodyWithJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := task.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idParam := ctx.Param("workspaceId")
	id, _ := strconv.Atoi(idParam)

	workspace := c.workspaceService.GetById(uint(id))
	if workspace.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "workspace not found",
		})
		return
	}

	task.Workspace = *workspace

	c.taskService.Create(&task)
	ctx.JSON(http.StatusCreated, task)
}
