package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kenzohfs/m/src/databases"
	"github.com/kenzohfs/m/src/migrations"
	"github.com/kenzohfs/m/src/routes"
)

func main() {
	db := databases.Connect()
	migrations.Migrate(db)

	r := gin.Default()
	routes.HandleRequests(r, db)

	r.Run()
}
