package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/mishaRomanov/test-ozon/config"
	"github.com/mishaRomanov/test-ozon/internal/handler"
	"github.com/mishaRomanov/test-ozon/internal/storage"
	cache "github.com/mishaRomanov/test-ozon/internal/storage/cache"
	"github.com/mishaRomanov/test-ozon/internal/storage/postgres"
	"github.com/sirupsen/logrus"
)

// func that creates storage depending on a flag value
func createStorageBasedOnFlag(config string, db *sql.DB) storage.Storager {
	if config == "postgres" {
		logrus.Infoln("Creating Postgres database...")
		return postgres.Create(db)
	}
	if config == "cache" {
		logrus.Infoln("Creating in-memory storage...")
		db.Close()
		return cache.NewCache()
	}
	logrus.Infoln("No correct config option detected. Creating in-memory storage...")
	return cache.NewCache()
}

func main() {
	//create a server
	service := gin.Default()

	//setting up config
	cfg, err := config.LoadConfig(".")

	connectString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Adress, cfg.DatabaseName)

	var database *sql.DB
	database, err = sql.Open("postgres", connectString)
	if err != nil {
		logrus.Errorf("Failed to open database: %v", err)
	}
	//creating a storage based on which flag value we got
	handlerObject := handler.New(createStorageBasedOnFlag(cfg.StorageType, database))
	//endpoint returns the full link if found
	//the short one is given through :shortLink parameter
	service.GET("/link/:shortLink", handlerObject.HandleGet)

	//endpoint creates and returns a new shortened link
	//the original one is sent through json
	service.POST("/link/add", handlerObject.HandlePost)

	//start listening
	logrus.Fatalf("%v", service.Run(":80"))

}
