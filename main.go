package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/golang-events-planning-backend/db"
	"github.com/golang-events-planning-backend/routes"
	"os"
)

func main() {
	migrate := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	db.InitDb()

	if *migrate {
		db.RunMigrations()
		os.Exit(0)
	}

	server := gin.Default()
	routes.RegisterEventRoutes(server)
	routes.RegisterUserRoutes(server)
	routes.RegisterCommentRoutes(server)

	server.Run(":3100")
}
