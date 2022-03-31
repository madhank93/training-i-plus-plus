package main

import (
	"ticket_reservation_system/config"
	"ticket_reservation_system/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.DatabaseConn()

	routes.UserRoutes(router)
	routes.TrainRoutes(router)

	router.Run(":5001")
}
