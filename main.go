package main

import (
	"VedioConverter/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	routes.SetupRouters(router)
	router.Run(":6000")

}
