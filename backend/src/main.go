package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:9000/"},
	}))

	router.GET("/itemsList", getItemsList)

	router.Run("localhost:8081")

}
