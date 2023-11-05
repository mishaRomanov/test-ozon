package shorten

import (
	"encoding/base64"
	"github.com/google/uuid"
)

func cleanShortLink(link string) string {
	//создаем закодированный в б64 uuid
	res := base64.StdEncoding.EncodeToString([]byte(link[:20]))
	return res[:10]
}

func MakeAShortLink(url string) string {
	//создаем новый рандомный uuid
	new := uuid.New()
	return cleanShortLink(new.String())

}
