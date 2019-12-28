package app

import (
	u "go-jwt-api/utils"
	"net/http"
)

// NotFoundHandler return error message
var NotFoundHandler = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "Resources not found in server"))
		next.ServeHTTP(w, r)
	})
}
