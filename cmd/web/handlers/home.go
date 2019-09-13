package handlers

import (
	"net/http"
)

// HomeIndex handler
func (g *HandlerContext) HomeIndex() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jsonData := []byte(`{"boom": "headshot"}`)
		w.Write(jsonData)
	})
}
