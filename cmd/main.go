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
	server := http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Addr:         ":8080",
	}

	//делаем обработчики
	router := mux.NewRouter()

	router.HandleFunc("/link/{link}", handler.HandleGet)

	router.HandleFunc("/link/test", handler.HandlePost)

	//врубаем
	logrus.Fatalf("%v", server.ListenAndServe())

}
