package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/Dadard29/go-api-net/models"
	"github.com/Dadard29/go-api-utils/log"
	"github.com/Dadard29/go-api-utils/log/logLevel"
	"io/ioutil"
	"net/http"
	"os"
)

var logger = log.NewLogger("REPOSITORY", logLevel.DEBUG)
var key = os.Getenv("API_KEY")

func LocationRepository(ip string, config map[string]string) models.Location {

	var location models.Location
	client := &http.Client{}

	urlWithParams := fmt.Sprintf(config["url"], key, ip)
	req, err := http.NewRequest(http.MethodGet, urlWithParams, nil)
	logger.CheckErr(err)

	resp, err := client.Do(req)
	logger.CheckErr(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &location)
	if err != nil {
		panic(err)
	}

	return location
}
