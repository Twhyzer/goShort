package util

import (
	"fmt"
	"math/rand"
	"net/url"
	"os"
)

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func ValidateDomain(domain string) bool {
	url, err := url.Parse(domain)

	if err != nil {
		return false
	}

	if url.Scheme != "http" && url.Scheme != "https" {
		return false
	}

	return true
}

func GenerateRouteKey(length int) string {
	bytes := make([]byte, length)

	for i := range bytes {
		bytes[i] = characters[rand.Intn(len(characters))]
	}

	return string(bytes)
}

func CreateShortDomain(routeKey string) string {
	domain := fmt.Sprintf("%s%s", os.Getenv("TARGET_DOMAIN"),routeKey)
	return domain
}
