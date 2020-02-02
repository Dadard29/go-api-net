package managers

import (
	. "github.com/Dadard29/go-api-net/api"
	"github.com/Dadard29/go-api-net/models"
	"github.com/Dadard29/go-api-net/repositories"
)

func LocationManager(ip string) models.Location {
	var config, _ = Api.Config.GetSubcategoryFromFile("iplocation", "api")

	return repositories.LocationRepository(ip, config)
}
