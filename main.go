package main

import (
	"ForumProject/model"
	"ForumProject/model/entity"
	"ForumProject/model/handler"
	"ForumProject/model/repository"
	"ForumProject/model/service"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const (
	config_path = "model/config.yaml"
)

func main() {

	var config *model.Config = loadConfig()
	db, err := model.NewDataBase(&config.DBConfig)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Post{}, &entity.Comment{})
	if err != nil {
		log.Fatalf("Failed to auto migrate database: %v", err)
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(*repositories)
	handlers := handler.NewHandler(services, db)
	router := handlers.InitRoutes()

	router.Run(":" + config.Port)
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
