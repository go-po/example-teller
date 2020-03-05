package api

import (
	"net/http"
	"strings"
)

func SplitPath(request *http.Request) (string, *http.Request) {
	if request.RequestURI == "/" {
		return "", request
	}
	i := strings.Index(request.RequestURI[1:], "/")
	if i < 0 {
		prefix := request.RequestURI[1:]
		request.RequestURI = "/"
		return prefix, request
	}
	prefix := request.RequestURI[1 : i+1]
	request.RequestURI = request.RequestURI[i+1:]
	return prefix, request
}
