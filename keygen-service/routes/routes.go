package routes

import (
	"net/http"

	"github.com/trite8q1/go-license-v2/keygen-service/routes/rest"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		rest.Index,
	},
	Route{
		"Health",
		"GET",
		"/health",
		rest.Health,
	},
	Route{
		"GenerateLicense",
		"GET",
		"/licenses",
		rest.GetLicenses,
	},
}
