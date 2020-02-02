package tests

import (
	"github.com/Dadard29/go-api-net/repositories"
	"testing"
)

func TestApi(t *testing.T) {
	ip := "51.83.46.84"

	config := map[string]string{
		"url": "https://api.ipgeolocation.io/ipgeo?apiKey=%s&ip=%s",
	}


	repositories.LocationRepository(ip, config)
}
