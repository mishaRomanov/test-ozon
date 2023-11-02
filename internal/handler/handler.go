package handler

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	link := mux.Vars(r)
	finalLink := link["link"]
	logrus.Println(finalLink)
	w.Write([]byte("handleGet works"))
	return
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Infof("Body reading error: &v", err)
	}
	defer r.Body.Close()
	logrus.Println(string(data))
	w.Write([]byte("it works"))
}
