package routes

import (
	controllers "github.com/Mardiniii/serapis/api/controllers"
)

// Routes collection
type Routes []Route

var routes = Routes{
	Route{
		Name:        "HealthCheck",
		Method:      "GET",
		Pattern:     "/api/healthcheck",
		HandlerFunc: controllers.HealthCheck,
	},
	Route{
		Name:        "Index Users",
		Method:      "GET",
		Pattern:     "/api/users",
		HandlerFunc: controllers.GetUsers,
	},
	Route{
		Name:        "Create User",
		Method:      "POST",
		Pattern:     "/api/users",
		HandlerFunc: controllers.CreateUser,
	},
	Route{
		Name:        "Get API Key",
		Method:      "GET",
		Pattern:     "/api/users/api-key",
		HandlerFunc: controllers.GetAPIKey,
	},
	Route{
		Name:        "Evaluation",
		Method:      "POST",
		Pattern:     "/api/evaluations/{platform}",
		HandlerFunc: controllers.CreateEvaluation,
	},
}
