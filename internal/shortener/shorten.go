package shorten

import (
	"encoding/base64"
	"github.com/google/uuid"
	storeErr "github.com/mishaRomanov/test-ozon/internal/storage"
	storage "github.com/mishaRomanov/test-ozon/internal/storage/cache"
	"github.com/sirupsen/logrus"
)

func cleanShortLink(link string) string {
	//создаем закодированный в б64 uuid
	res := base64.StdEncoding.EncodeToString([]byte(link[:20]))
	return res[:10]
}

func MakeAShortLink(url string, inmemory *storage.Cache) (string, error) {
	//проверяем, есть ли такая ссылка в мапе
	for key, link := range inmemory.Cache {
		if link == url {
			logrus.Errorf("An attempt to create a short link: %v: \n%s", storeErr.ErrAlreadyExists, key)
			return "", storeErr.ErrAlreadyExists
		}
	}
	//создаем новый рандомный uuid
	new := uuid.New()
	return cleanShortLink(new.String()), nil
}
