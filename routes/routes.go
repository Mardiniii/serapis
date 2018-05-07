package routes

import (
	controller "../controllers"
)

// Routes collection
type Routes []Route

var routes = Routes{
	Route{
		Name:        "HealthCheck",
		Method:      "GET",
		Pattern:     "/healthcheck",
		HandlerFunc: controller.HealthCheck,
	},
	Route{
		Name:        "Index Users",
		Method:      "GET",
		Pattern:     "/users",
		HandlerFunc: controller.UserIndex,
	},
	Route{
		Name:        "Create User",
		Method:      "POST",
		Pattern:     "/users",
		HandlerFunc: controller.UserCreate,
	},
}
