package controllers

import (
	"github.com/Dadard29/go-api-utils/log"
	"github.com/Dadard29/go-api-utils/log/logLevel"
	"io/ioutil"
	"net/http"
)

var logger = log.NewLogger("CONTROLLER", logLevel.DEBUG)

func SwaggerController(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	dat, err := ioutil.ReadFile("assets/swagger.json")
	logger.CheckErr(err)
	_, err = w.Write(dat)
	logger.CheckErr(err)
}

