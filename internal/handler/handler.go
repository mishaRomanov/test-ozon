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

// структура для парсинга json
type requestBody struct {
	Url string `json:"url"`
}

// здесь создаем in-memory хранилище для значений
// надо сделать универсальную функцию которая создает нужное в зависимости
// от переданного флага для запуска значение (мапа или бд)
var store = storage.NewCache()

// хендлит гет реквесты по типу localhost:8080/link/*
func HandleGet(ctx *gin.Context) {
	//достаем параметр (сокращенную ссылку)
	shortLink := ctx.Param("shortLink")

	//проверяем, не пусто ли
	if shortLink == "" {
		ctx.String(http.StatusBadRequest, "Empty link")
		return
	}
	//находим полную ссылку-пару к короткой ссылке
	redirectTo, err := store.GetValue(shortLink)
	logrus.Infoln(redirectTo)
	if err != nil {
		//проверяем ошибку
		logrus.Errorf("Error while searching for value %v", err)
		ctx.String(fiber.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, redirectTo)
}

// хендлер для создания ссылки
func HandlePost(ctx *gin.Context) {
	body := requestBody{}
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.String(fiber.StatusBadRequest, "Invalid JSON")
		return
	}
	//записываем старую ссылку и новую
	oldUrl := body.Url
	newUrl := shorten.MakeAShortLink(oldUrl)

	//записываем значения
	err = store.WriteValue(newUrl, oldUrl)
	if err != nil {
		logrus.Errorf("%v", err)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	logrus.Infof("Data written: old link - %s, new link - %s\n", oldUrl, newUrl)
	ctx.String(http.StatusOK, fmt.Sprintf("New link generated: localhost:80/link/%s", newUrl))
}
