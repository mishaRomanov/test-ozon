package shorten

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func MakeAShortLink(url string) string {
	shortUrl := "shorti.fy/"
	sum := sha256.Sum256([]byte(url))
	b64 := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s", sum[:10])))
	shortUrl += b64[:10]
	shortUrl = strings.TrimRight(shortUrl, "=")
	finalUrl := shortUrl[:len(shortUrl)/2] + "_" + shortUrl[len(shortUrl)/2+1:]
	return finalUrl
}
