package handler

import (
	"github.com/gorilla/mux"
	storage "github.com/mishaRomanov/test-ozon/internal/storage/cache"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"

	"github.com/mishaRomanov/test-ozon/internal/shortener"
)

// здесь создаем in-memory хранилище для значений
var store = storage.NewCache()

// хендлит гет реквесты по типу localhost:8080/link/"link"
func HandleGet(w http.ResponseWriter, r *http.Request) {
	link := mux.Vars(r)
	finalLink, ok := link["link"]
	logrus.Infof("New request to redirect to: %s\n", finalLink)
	//проверяем, нашлось ли значение в пути URL и проверяем, не пустое ли оно
	if !ok || finalLink == "" {
		logrus.Error("Parameter not found")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Parameter not found"))
		return
	}
	redirectTo, err := store.GetValue(finalLink)
	if err != nil {
		logrus.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	http.Redirect(w, r, redirectTo, 200)

}

// хендлер для создания ссылки
func HandlePost(w http.ResponseWriter, r *http.Request) {
	//читаем тело запроса
	data, err := io.ReadAll(r.Body)
	//проверяем ошибку чтения
	if err != nil {
		logrus.Infof("Body reading error: &v", err)
	}
	//деферим закрытие чтения тела
	defer r.Body.Close()

	oldUrl := string(data)
	newUrl := shorten.MakeAShortLink(oldUrl)
	//записываем значения
	err = store.WriteValue(newUrl, oldUrl)
	if err != nil {
		logrus.Errorf("%v", err)
	}
	logrus.Infof("%v", store)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(newUrl))
}
