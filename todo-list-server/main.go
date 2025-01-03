package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kenzohfs/m/src/databases"
	"github.com/kenzohfs/m/src/routes"
)

func main() {
	db := databases.Connect()

	r := gin.Default()
	routes.HandleRequests(r, db)

	r.Run()
}
