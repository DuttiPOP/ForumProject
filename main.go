package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	router.POST("/register", func(context *gin.Context) {

	})

	router.DELETE("/user/:id", func(context *gin.Context) {

	})

	router.PATCH("/user/:id", func(context *gin.Context) {

	})

}
