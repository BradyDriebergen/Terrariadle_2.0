package api

import "net/http"

// Not currently in use. Used to use cors when testing against frontend because they were
// different ports. The frontend now uses a vite proxy to hit the backend so this isn't
// need anymore. Re-implement if you ever make a mobile app or another project that needs
// to access the backend from a different origin.

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
