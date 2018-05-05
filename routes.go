package main

import (
	"net/http"
)

// Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes collection
type Routes []Route

var routes = Routes{
	Route{
		Name:        "HealthCheck",
		Method:      "GET",
		Pattern:     "/healthcheck",
		HandlerFunc: HealthCheck,
	},
	Route{
		Name:        "Index Users",
		Method:      "GET",
		Pattern:     "/users",
		HandlerFunc: UserIndex,
	},
	Route{
		Name:        "Create User",
		Method:      "POST",
		Pattern:     "/users",
		HandlerFunc: UserCreate,
	},
}
