package util

import (
	"fmt"
	"math/rand"
	"net/url"
	"os"
)

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Validator for URLs
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

// Generates a random string with a specified length
func RandomString(length int) string {
	bytes := make([]byte, length)

	for i := range bytes {
		bytes[i] = characters[rand.Intn(len(characters))]
	}

	return string(bytes)
}

// Creates the short URL from the target domain and the short key
func CreateShortDomain(routeKey string) string {
	domain := fmt.Sprintf("%s%s", os.Getenv("TARGET_DOMAIN"), routeKey)
	return domain
}
