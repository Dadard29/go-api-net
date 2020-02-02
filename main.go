package main

import (
	. "github.com/Dadard29/go-api-net/api"
	"github.com/Dadard29/go-api-net/controllers"
	"github.com/Dadard29/go-api-utils/API"
	"github.com/Dadard29/go-api-utils/service"
	"net/http"
)

//var routes = map[string]func(w http.ResponseWriter, r *http.Request) {
var routes = service.RouteMapping{
	Mapping: map[string]service.Route{
		"/location": {Handler: controllers.LocationRoute, Method:  http.MethodGet},
		"/swagger": {Handler: controllers.SwaggerController, Method:  http.MethodGet},
	},
}


func main() {

	Api = API.NewAPI("Api-Net", "config/config.json", routes, true)

	Api.Service.Start()

	Api.Service.Stop()
}
