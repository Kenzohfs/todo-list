package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenzohfs/m/src/models"
	"github.com/kenzohfs/m/src/services"
)

type TaskController struct {
	taskService *services.TaskService
}

func NewTaskController(taskService *services.TaskService) *TaskController {
	return &TaskController{taskService: taskService}
}

func (c *TaskController) Create(ctx *gin.Context) {
	var task models.Task

	if err := ctx.ShouldBindBodyWithJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.taskService.Create(&task)
	ctx.JSON(http.StatusCreated, task)
}
