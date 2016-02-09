package handlers

import (
	"net/http"
)

// HTTPMethod represents the possible HTTP methods
type HTTPMethod string

const (
	HTTPGet     HTTPMethod = "GET"
	HTTPPost    HTTPMethod = "POST"
	HTTPPut     HTTPMethod = "PUT"
	HTTPDelete  HTTPMethod = "DELETE"
	HTTPHead    HTTPMethod = "HEAD"
	HTTPPatch   HTTPMethod = "PATCH"
	HTTPOptions HTTPMethod = "OPTIONS"
)

// String is the fmt.Stringer interface implementation
func (m HTTPMethod) String() string {
	return string(m)
}

// Handler is the interface to represent a http.Handler, plus information about the path and HTTP method that it serves
type Handler interface {
	http.Handler
	PathInfo() (HTTPMethod, string)
}
