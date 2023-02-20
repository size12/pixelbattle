package middleware

import "net/http"

func OnlyJSONContent(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType != "application/json" {
			http.Error(w, "Wrong content type. Accept only application/json", http.StatusBadRequest)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
