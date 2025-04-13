package openapi

import "net/http"

func MaxBytes(inner http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 5*1000000)
		inner(w, r)
	}
}
