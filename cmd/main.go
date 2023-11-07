package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/mishaRomanov/test-ozon/internal/handler"
	"github.com/mishaRomanov/test-ozon/internal/storage"
	cache "github.com/mishaRomanov/test-ozon/internal/storage/cache"
	"github.com/mishaRomanov/test-ozon/internal/storage/postgres"
	"github.com/sirupsen/logrus"
)

// func that creates storage depending on a flag value
func createStorageBasedOnFlag(flag *string) storage.Storager {
	if *flag == "postgres" {
		return postgres.NewDatabase()
	}
	if *flag == "cache" {
		return cache.NewCache()
	}
	logrus.Infoln("Wrong flag value given. Creating a cache storage...")
	return nil
}

func main() {
	//create a server
	service := gin.Default()
	Type := flag.String("storage-type", "cache", "Used to determine what kind of storage to use")
	flag.Parse()

	//creating a storage based on which flag value we got
	handlerObject := handler.New(createStorageBasedOnFlag(Type))
	//endpoint returns the full link if found
	//the short one is given through :shortLink parameter
	service.GET("/link/:shortLink", handlerObject.HandleGet)

	//endpoint creates and returns a new shortened link
	//the original one is sent through json
	service.POST("/link/add", handlerObject.HandlePost)

	//start listening
	logrus.Fatalf("%v", service.Run(":80"))

}
