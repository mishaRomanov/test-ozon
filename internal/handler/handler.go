package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/mishaRomanov/test-ozon/internal/shortener"
	storage "github.com/mishaRomanov/test-ozon/internal/storage/cache"
	"github.com/sirupsen/logrus"

	"net/http"
)

// json parsing struct
type requestBody struct {
	Url string `json:"url"`
}

//creating a new storage
var store = storage.NewCache()

//GET requests handler
func HandleGet(ctx *gin.Context) {
	//extract a parameter
	shortLink := ctx.Param("shortLink")

	//check for " "
	if shortLink == "" {
		ctx.String(http.StatusBadRequest, "Empty link")
		return
	}
	//search for a pair
	redirectTo, err := store.GetValue(shortLink)
	logrus.Infoln(redirectTo)
	if err != nil {
		//handling error
		logrus.Errorf("Error while searching for value %v", err)
		ctx.String(fiber.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, redirectTo)
}

//POST requests handler
func HandlePost(ctx *gin.Context) {
	body := requestBody{}
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.String(fiber.StatusBadRequest, "Invalid JSON")
		return
	}
	//writing old and new links
	oldUrl := body.Url
	newUrl, err := shorten.MakeAShortLink(oldUrl, store)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	//writing links into storage
	err = store.WriteValue(newUrl, oldUrl)
	if err != nil {
		logrus.Errorf("%v", err)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	logrus.Infof("Data written: old link - %s, new link - %s\n", oldUrl, newUrl)
	ctx.String(http.StatusOK, fmt.Sprintf("New link generated: localhost:80/link/%s", newUrl))
}
