package hello

import (
	"net/http"

	"github.com/starius/api2"
)

func GetRoutes(s *EchoService) []api2.Route {
	return []api2.Route{
		{Method: http.MethodGet, Path: "/hello", Handler: api2.Method(&s, "Hello")},
	}
}
