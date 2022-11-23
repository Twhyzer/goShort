package shorter

import (
	"errors"

	"github.com/Twhyzer/goShort/src/internal/database"
	"github.com/Twhyzer/goShort/src/internal/util"
)

const routKeyLength = 5

// Creates a shorturl after checking the targetUrl and adds it to the database
func CreateShortURL(domain string) (string, error) {

	if !util.ValidateDomain(domain) {
		return "", errors.New("DOMAIN IS NOT VALID")
	}

	requestKey := GenerateRouteKey()

	res, err := database.InsertShortURL(domain, requestKey)

	if err != nil {
		return "", err
	}

	return res, nil
}

// Recursive function. Generates a previously non-existent key.
func GetShortByKey(domain, key string) (database.Shorts, error) {

	short, err := database.GetShortByKey(key)

	if err != nil {
		return database.Shorts{}, nil
	}

	return short, nil
}

// Generates a route key that does not exist
func GenerateRouteKey() string {
	routeKey := util.RandomString(routKeyLength)
	_, err := database.GetShortByKey(routeKey)

	if err != nil {
		return routeKey
	}

	return GenerateRouteKey()
}
