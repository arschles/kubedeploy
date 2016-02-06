package handlers

import (
	"net/http"
)

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

func (m HTTPMethod) String() string {
	return string(m)
}

type Handler interface {
	http.Handler
	PathInfo() (HTTPMethod, string)
}
