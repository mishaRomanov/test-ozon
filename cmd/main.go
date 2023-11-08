package main

import (
	"database/sql"
	"flag"
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
func createStorageBasedOnFlag(flag *string, db *sql.DB) storage.Storager {
	if *flag == "postgres" {
		return postgres.Create(db)
	}
	if *flag == "cache" {
		db.Close()
		return cache.NewCache()
	}
	logrus.Infoln("Wrong flag value given. Creating a cache storage...")
	return cache.NewCache()
}

func main() {
	//create a server
	service := gin.Default()
	//creating a flag
	Type := flag.String("storage", "cache", "Used to determine what kind of storage to use")
	//parsing
	flag.Parse()

	//setting up config
	cfg, err := config.LoadConfig("../config")

	connectString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Adress, cfg.DatabaseName)

	var database *sql.DB
	database, err = sql.Open("postgres", connectString)
	if err != nil {
		logrus.Errorf("Failed to open database: %v", err)
	}
	//creating a storage based on which flag value we got
	handlerObject := handler.New(createStorageBasedOnFlag(Type, database))
	//endpoint returns the full link if found
	//the short one is given through :shortLink parameter
	service.GET("/link/:shortLink", handlerObject.HandleGet)

	//endpoint creates and returns a new shortened link
	//the original one is sent through json
	service.POST("/link/add", handlerObject.HandlePost)

	//start listening
	logrus.Fatalf("%v", service.Run(":80"))

}
