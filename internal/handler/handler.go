package handler

import (
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("handleGet works"))
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Infof("Body reading error: &v", err)
	}

}
