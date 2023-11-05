package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mishaRomanov/test-ozon/internal/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	//создаем сервер
	service := gin.Default()

	//эндпоинт возвращает оригинальную ссылку, в пути указываем сокращенную
	service.GET("/link/:shortLink", handler.HandleGet)

	//эндпоинт создает короткую ссылку из оригинальной ссылки отправленной в json
	service.POST("/link/", handler.HandlePost)

	//запускаем сервер
	logrus.Fatalf("%v", service.Run(":80"))

}
