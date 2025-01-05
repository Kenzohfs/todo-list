package controllers

import (
	"net/http"
	"strconv"

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

func (c *WorkspaceController) Get(ctx *gin.Context) {
	pageQuery := ctx.Query("page")
	perPageQuery := ctx.Query("perPage")
	sortBy := ctx.Query("sortBy")
	sortDirection := ctx.Query("sortDirection")
	filter := ctx.Query("filter")

	page, _ := strconv.Atoi(pageQuery)
	perPage, _ := strconv.Atoi(perPageQuery)

	result := c.workspaceService.Get(page, perPage, sortBy, sortDirection, filter)

	ctx.JSON(http.StatusOK, result)
}

func (c *WorkspaceController) GetById(ctx *gin.Context) {
	idParam := ctx.Param("workspaceId")

	id, _ := strconv.Atoi(idParam)

	workspace := c.workspaceService.GetById(uint(id))

	if workspace.ID == 0 {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	ctx.JSON(http.StatusOK, workspace)
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
