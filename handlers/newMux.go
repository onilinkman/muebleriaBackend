package handlers

import (
	"fmt"
	"net/http"
)

type customeHandler func(w http.ResponseWriter, r *http.Request)

func HeadersMiddleware(h http.Handler, method string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		fmt.Println("hola")
		if r.Method != method {
			http.Error(w, "Method not permitted", http.StatusMethodNotAllowed)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow origin
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Allow headers
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")

		// Handle preflight request
		if r.Method == "OPTIONS" {
			return
		}
		fmt.Println("hola")

		// Pass request to the next handler
		next.ServeHTTP(w, r)
	})
}

func AddHeaders(f func(w http.ResponseWriter, r *http.Request), method string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", method)
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

		if r.Method != method {
			http.Error(w, "Method not permitted", http.StatusMethodNotAllowed)
			return
		}
		f(w, r)
	})

}
