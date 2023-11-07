package shorten

import (
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/mishaRomanov/test-ozon/internal/storage"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"unicode"
)

// we use that func in case it has the same index as slash
func generateNum(n int) int {
	return rand.Intn(n)
}

// this func checks whether the link has numbers in it or not
func cleanShortLink(link string) string {
	slash := rand.Intn(len(link))
	link = link[:slash] + "_" + link[slash+1:]
	for {
		for _, char := range link {
			//if we have at least one num. return
			if unicode.IsNumber(char) {
				return link
			}
		}
		//generating a random num
		num := rand.Intn(len(link))
		//checking if slash index is equal to num
		if num == slash {
			num = generateNum(len(link))
			link = link[:num] + strconv.Itoa(num) + link[num+1:]
			break
		}
		link = link[:num] + strconv.Itoa(num) + link[num+1:]
		break
	}
	return link
}

func MakeAShortLink(url string, dataStorage storage.Storager) (string, error) {
	//checking if we already have that url
	logrus.Infoln(url)
	ok, err := dataStorage.LookUp(url)
	if ok {
		logrus.Errorf("An attempt to create a short link: %v: \n", storage.ErrAlreadyExists)
		return "", storage.ErrAlreadyExists
	}
	if err != nil {
		logrus.Errorf("%v", err)
		return "", err
	}
	//creating UUID
	new := uuid.New()
	//Encode uuid to base64
	res := base64.StdEncoding.EncodeToString([]byte(new.String()[:20]))
	//we shorten uuid and then use only 10 elems of string
	//in order to get rid of b64 junk like "==="
	res = res[:10]
	return cleanShortLink(res), nil
}
