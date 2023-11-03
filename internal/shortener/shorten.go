package shorten

import "crypto/sha256"

func makeAShortLink(url string) string {
	shortUrl := "shorti.fy/"
	sum := sha256.Sum256([]byte(url))

}
