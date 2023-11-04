package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"

	"github.com/mishaRomanov/test-ozon/internal/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	//создаем сервер
	router := mux.NewRouter()
	server := http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Addr:         ":8080",
		Handler:      router,
	}

	//делаем обработчики
	router.HandleFunc("/link/{link}", handler.HandleGet)

	router.HandleFunc("/link/", handler.HandlePost)

	logrus.Infoln("Starting the service!")

	//запускаем сервер
	logrus.Fatalf("%v", server.ListenAndServe())

}
