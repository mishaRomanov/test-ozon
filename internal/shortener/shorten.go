package shorten

import (
	"math/rand"
	"strings"

	"github.com/mishaRomanov/test-ozon/internal/storage"
	"github.com/sirupsen/logrus"
)

func MakeAShortLink(url string, dataStorage storage.Storager) (string, error) {
	//checking if we already have that url
	ok, err := dataStorage.LookUp(url)
	if ok {
		logrus.Errorf("An attempt to create a short link: %v: \n", storage.ErrAlreadyExists)
		return "", storage.ErrAlreadyExists
	}
	if err != nil {
		logrus.Errorf("%v", err)
		return "", err
	}
	//creating a string with all needed characters
	var alphabet = "ynAJfoSgdXHB5VasEMtcbPCr1uNZ4LG723ehWkvwYR6KpxjTm8iQUFqz9D"

	//let's write len of the string into a variable
	var alphabetLen = len(alphabet)
	//creating a builder for building a string
	var builder strings.Builder
	//writing "_" cause we need it
	builder.WriteString("_")

	for i := 0; i < 9; i++ {
		builder.WriteString(string(alphabet[rand.Intn(alphabetLen)]))
	}
	return builder.String(), nil

}
