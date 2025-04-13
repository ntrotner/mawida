package openapi

import (
	"net/http"
	"strings"
)

func ReadTokenFromHeader(r *http.Request) (string, bool) {
	authorizationHeader := r.Header.Get("Authorization")
	_, after, found := strings.Cut(authorizationHeader, "Bearer ")
	return after, found
}

func WriteTokenToHeader(token *string, w http.ResponseWriter) {
	w.Header().Add("Authorization", "Bearer "+*token)
}
