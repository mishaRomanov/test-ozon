package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mishaRomanov/test-ozon/internal/shortener"
	"github.com/mishaRomanov/test-ozon/internal/storage"
	"github.com/sirupsen/logrus"
	"net/http"
)

// json parsing struct
type requestBody struct {
	Url string `json:"url"`
}

// Handler struct that contains storage inside
// Independent of storage type
type Handler struct {
	DataStorage storage.Storager
}

// GET requests handler method
func (h *Handler) HandleGet(ctx *gin.Context) {
	//extract a parameter
	shortLink := ctx.Param("shortLink")
	//check for " "
	if shortLink == "" {
		ctx.String(http.StatusBadRequest, "Empty link")
		return
	}
	//search for a pair
	redirectTo, err := h.DataStorage.GetValue(shortLink)
	if err != nil {
		//handling error
		logrus.Errorf("Error while searching for value %v", err)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, redirectTo)
}

// POST requests handler
func (h *Handler) HandlePost(ctx *gin.Context) {
	body := requestBody{}
	err := ctx.BindJSON(&body)
	if err != nil {
		logrus.Errorf("Invalid JSON: %v", err)
		ctx.String(http.StatusBadRequest, "Invalid JSON")
		return
	}
	//if body if nil return badrequest
	if body.Url == "" {
		ctx.String(http.StatusBadRequest, "No body given")
		return
	}
	//writing old and new links
	oldUrl := body.Url
	newUrl, err := shorten.MakeAShortLink(oldUrl, h.DataStorage)
	if err != nil {
		logrus.Errorf("%v", err)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	//writing links into storage
	err = h.DataStorage.WriteValue(newUrl, oldUrl)
	if err != nil {
		logrus.Errorf("%v", err)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	logrus.Infof("Data written: old link - %s, new link - %s\n", oldUrl, newUrl)
	ctx.String(http.StatusOK, fmt.Sprintf("New link generated: localhost:8080/link/%s", newUrl))
}

func New(storager storage.Storager) *Handler {
	object := Handler{
		storager,
	}
	return &object
}
