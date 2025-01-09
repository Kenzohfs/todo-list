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

func (c *TaskController) Get(ctx *gin.Context) {
	idParam := ctx.Param("workspaceId")
	id, _ := strconv.Atoi(idParam)

	workspace := c.workspaceService.GetById(uint(id))
	if workspace.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "workspace not found",
		})
		return
	}

	pageQuery := ctx.Query("page")
	perPageQuery := ctx.Query("perPage")
	sortBy := ctx.Query("sortBy")
	sortDirection := ctx.Query("sortDirection")
	filter := ctx.Query("filter")
	status := ctx.Query("status")

	page, _ := strconv.Atoi(pageQuery)
	perPage, _ := strconv.Atoi(perPageQuery)

	result := c.taskService.Get(id, page, perPage, sortBy, sortDirection, filter, status)

	ctx.JSON(http.StatusOK, result)
}

func (c *TaskController) Edit(ctx *gin.Context) {
	taskIdParam := ctx.Param("taskId")
	taskId, _ := strconv.Atoi(taskIdParam)

	task := c.taskService.GetById(uint(taskId))

	if task.ID == 0 {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	if err := ctx.ShouldBindJSON(&task); err != nil {
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

	c.taskService.Edit(task)
	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) Delete(ctx *gin.Context) {
	workspaceIdParam := ctx.Param("workspaceId")
	workspaceId, _ := strconv.Atoi(workspaceIdParam)

	workspace := c.workspaceService.GetById(uint(workspaceId))
	if workspace.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "workspace not found",
		})
		return
	}

	taskIdParam := ctx.Param("taskId")
	taskId, _ := strconv.Atoi(taskIdParam)

	c.taskService.Delete(uint(taskId))
	ctx.JSON(http.StatusNoContent, nil)
}
