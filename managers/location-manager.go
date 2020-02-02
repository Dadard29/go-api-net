package managers

import (
	"github.com/Dadard29/go-api-net/models"
	"github.com/Dadard29/go-api-net/repositories"
)

func LocationManager(ip string) models.Location {
	return repositories.LocationRepository(ip)
}
