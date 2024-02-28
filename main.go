package main

import (
	"ForumProject/model"
	"ForumProject/model/entity"
	"ForumProject/model/repository"
	"ForumProject/model/service"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
)

const (
	config_path = "model/config.yaml"
)

func main() {

	router := gin.Default()

	var config *model.Config = loadConfig()
	db, err := model.NewDataBase(&config.DBConfig)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	repository := repository.NewRepository(db)
	service := service.NewService(*repository)

	router.POST("/register", func(context *gin.Context) {
		var user entity.User
		err := context.BindJSON(&user)
		if err != nil {
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		userID, err := service.IUserService.Create(user)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"id": userID})

	})
	router.GET("/user/:id", func(context *gin.Context) {
		user, err := service.IUserService.Get(context.Params.ByName("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, user)
	})

	router.DELETE("/user/:id", func(context *gin.Context) {
		err := service.IUserService.Delete(context.Params.ByName("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, "{}")
	})

	router.PATCH("/user/:id", func(context *gin.Context) {

	})

	router.Run(":8080")
}

func loadConfig() *model.Config {
	yamlFile, err := os.ReadFile(config_path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	var config model.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &config
}
