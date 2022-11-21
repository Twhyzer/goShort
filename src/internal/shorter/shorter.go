package shorter

import (
	"errors"

	"github.com/Twhyzer/goShort/src/internal/database"
	"github.com/Twhyzer/goShort/src/internal/util"
)

const routKeyLength = 5

func CreateShortURL(domain string) (string, error) {

	if !util.ValidateDomain(domain) {
		return "", errors.New("DOMAIN IS NOT VALID")
	}

	routeKey := util.GenerateRouteKey(routKeyLength)
	
	res, err := database.InsertShortURL(domain, routeKey)

	if err != nil {
		return "", err;
	}

	return res, nil
}