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
		Pattern:     "/api/healthcheck",
		HandlerFunc: controllers.HealthCheck,
	},
	Route{
		Name:        "Index Users",
		Method:      "GET",
		Pattern:     "/api/users",
		HandlerFunc: controllers.UserIndex,
	},
	Route{
		Name:        "Create User",
		Method:      "POST",
		Pattern:     "/api/users",
		HandlerFunc: controllers.UserCreate,
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
		Pattern:     "/api/eval/:platform",
		HandlerFunc: controllers.CreateEvaluation,
	},
}
