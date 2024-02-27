package main

import (
	"ForumProject/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const (
	config_path = "model/config.yaml"
)

func main() {

	router := gin.Default()

	var dbconfig *model.DatabaseConfig = loadConfig()
	db, err := model.NewDataBase(*dbconfig)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	router.POST("/register", func(context *gin.Context) {

	})

	router.DELETE("/user/:id", func(context *gin.Context) {

	})

	router.PATCH("/user/:id", func(context *gin.Context) {

	})

}

func loadConfig() *model.DatabaseConfig {
	yamlFile, err := os.ReadFile(config_path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	var dbconfig model.DatabaseConfig
	err = yaml.Unmarshal(yamlFile, dbconfig)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &dbconfig
}
