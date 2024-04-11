package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chiennd172002/golang-swagger/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := routes.SetupV1Router()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	fmt.Println("Server is running on port", port)
	r.Run(":" + port)
}
