package controllers

import (
	"github.com/Dadard29/go-api-net/managers"
	"github.com/Dadard29/go-api-utils/API"
	"net"
	"net/http"
)

func LocationRoute(w http.ResponseWriter, r *http.Request) {
	var ip string

	ipAddr := r.Header.Get("X-FORWARDED-FOR")
	if ipAddr == "" {
		var err error
		ip, _, err = net.SplitHostPort(r.RemoteAddr)
		logger.CheckErr(err)
	} else {
		ip = ipAddr
	}

	logger.Debug(ip)

	res := managers.LocationManager(ip)
	logger.CheckErr(API.BuildJsonResponse(true, "test done", res, w))
}

