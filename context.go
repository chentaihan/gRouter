package gRouter

import (
	"net/http"
)

type Context struct {
	writermem responseWriter
	Request   *http.Request
}

type responseWriter struct {
	http.ResponseWriter
	size   int
	status int
}
