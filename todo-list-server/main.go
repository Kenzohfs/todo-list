package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kenzohfs/m/src/databases"
	"github.com/kenzohfs/m/src/migrations"
	"github.com/kenzohfs/m/src/routes"
)

func main() {
	db := databases.Connect()
	migrations.Migrate(db)

	r := gin.Default()
	r.Use(cors.Default())
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"X-Custom-Header"},
	// 	AllowCredentials: true,
	// }))

	routes.HandleRequests(r, db)

	r.Run()
}
