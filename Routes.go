package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:    "Index",
		Method:  "GET",
		Pattern: "/",
		Handler: Index,
	},
	Route{
		Name:    "TodoIndex",
		Method:  "GET",
		Pattern: "/todos",
		Handler: TodoIndex,
	},
	Route{
		Name:    "TodoShow",
		Method:  "GET",
		Pattern: "/todos/{todoId}",
		Handler: TodoShow,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.Handler)
	}

	return router
}
