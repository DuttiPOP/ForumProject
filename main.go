package main

import "github.com/gin-gonic/gin"
const (
	config_path = "model/config.yaml"
)

func main() {

	router := gin.Default()

	router.POST("/register", func(context *gin.Context) {

	})

	router.DELETE("/user/:id", func(context *gin.Context) {

	})

	router.PATCH("/user/:id", func(context *gin.Context) {

	})

}
