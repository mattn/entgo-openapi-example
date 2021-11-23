/*
 * Ent Schema API
 *
 * This is an auto generated API description made out of an Ent schema definition
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"CreateEntry",
		strings.ToUpper("Post"),
		"/entries",
		CreateEntry,
	},

	Route{
		"DeleteEntry",
		strings.ToUpper("Delete"),
		"/entries/{id}",
		DeleteEntry,
	},

	Route{
		"ListEntry",
		strings.ToUpper("Get"),
		"/entries",
		ListEntry,
	},

	Route{
		"ReadEntry",
		strings.ToUpper("Get"),
		"/entries/{id}",
		ReadEntry,
	},

	Route{
		"UpdateEntry",
		strings.ToUpper("Patch"),
		"/entries/{id}",
		UpdateEntry,
	},
}
