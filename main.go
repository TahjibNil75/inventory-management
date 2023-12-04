package main

import (
	"fmt"

	"github.com/TahjibNil75/inventory-management/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading godotenv: ", err)
	}

	r := gin.Default()
	routes.SetUpRoutes(r)
	r.Run(":8080")
}
