package shorten

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func makeAShortLink(url string) string {
	shortUrl := "shorti.fy/"
	sum := sha256.Sum256([]byte(url))
	b64 := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s", sum[:10])))
	shortUrl += b64[:10]

	shortUrl = strings.Replace(shortUrl, "=", "", 5)
	return shortUrl

}
