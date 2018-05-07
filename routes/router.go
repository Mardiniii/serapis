package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	middlewares "../middlewares"
)

// Router returns a new router with all the routes
func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = middlewares.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
