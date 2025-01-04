package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kenzohfs/m/src/controllers"
	"github.com/kenzohfs/m/src/repos"
	"github.com/kenzohfs/m/src/services"
	"gorm.io/gorm"
)

func HandleRequests(r *gin.Engine, db *gorm.DB) {
	workspaceRepo := repos.NewWorkspaceRepo(db)
	workspaceService := services.NewWorkspaceService(workspaceRepo)
	workspaceController := controllers.NewWorkspaceController(workspaceService)

	taskRepo := repos.NewTaskRepo(db)
	taskService := services.NewTaskService(taskRepo)
	taskController := controllers.NewTaskController(taskService)

	r.GET("/health", controllers.Health)

	r.POST("/api/v1/workspaces", workspaceController.Create)
	r.GET("/api/v1/workspaces", workspaceController.Get)
	r.GET("/api/v1/workspaces/:workspaceId", workspaceController.GetById)

	r.POST("/api/v1/workspaces/:workspaceId/tasks", taskController.Create)
}
