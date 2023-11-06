package shorten

import (
	"encoding/base64"
	"github.com/google/uuid"
	storeErr "github.com/mishaRomanov/test-ozon/internal/storage"
	storage "github.com/mishaRomanov/test-ozon/internal/storage/cache"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"unicode"
)

// функция выполняет проверку урл
// на наличие чисел
func cleanShortLink(link string) string {
	for {
		for _, char := range link {
			//если хотя бы один элемент это число то завершаем цикл и возвращаем строку целиком
			if unicode.IsNumber(char) {
				return link
			}
		}
		//генерим случайное число индекс и вставляем его в качестве строки
		num := rand.Intn(len(link))
		link = link[:num] + strconv.Itoa(num) + link[num+1:]
		break
	}
	return link
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
	//создаем закодированный в б64 uuid
	res := base64.StdEncoding.EncodeToString([]byte(new.String()[:20]))
	//делаем строку с первыми 10 элементами
	//это убирает из строки все лишние элементы после
	//кодирования в base64 (типа == и прочее)
	res = res[:10]
	//генерим случайное число и по его индексу вставляем слэш
	slash := rand.Intn(len(res))
	res = res[:slash] + "_" + res[slash+1:]
	return cleanShortLink(res), nil
}
