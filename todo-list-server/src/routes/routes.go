package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kenzohfs/m/src/controllers"
	"gorm.io/gorm"
)

func HandleRequests(r *gin.Engine, db *gorm.DB) {
	r.GET("/health", controllers.Health)
}
