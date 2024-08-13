package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-events-planning-backend/db"
	"github.com/golang-events-planning-backend/routes"
)

func main() {
	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":3100")
}
