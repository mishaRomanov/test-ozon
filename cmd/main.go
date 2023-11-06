package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mishaRomanov/test-ozon/internal/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	//create a server
	service := gin.Default()

	//endpoint returns the full link if found
	//the short one is given through :shortLink parameter
	service.GET("/link/:shortLink", handler.HandleGet)

	//endpoint creates and returns a new shortened link
	//the original one is sent through json
	service.POST("/link/add", handler.HandlePost)

	//start listening
	logrus.Fatalf("%v", service.Run(":80"))

}
