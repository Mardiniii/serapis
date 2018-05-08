package routes

import (
	controllers "github.com/Mardiniii/serapis_api/controllers"
)

// Routes collection
type Routes []Route

var routes = Routes{
	Route{
		Name:        "HealthCheck",
		Method:      "GET",
		Pattern:     "/healthcheck",
		HandlerFunc: controllers.HealthCheck,
	},
	Route{
		Name:        "Index Users",
		Method:      "GET",
		Pattern:     "/users",
		HandlerFunc: controllers.UserIndex,
	},
	Route{
		Name:        "Create User",
		Method:      "POST",
		Pattern:     "/users",
		HandlerFunc: controllers.UserCreate,
	},
	Route{
		Name:        "Get API Key",
		Method:      "GET",
		Pattern:     "/users/api-key",
		HandlerFunc: controllers.GetAPIKey,
	},
}
