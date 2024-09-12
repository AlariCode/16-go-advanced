package middleware

import (
	"fmt"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging")
		next.ServeHTTP(w, r)
		fmt.Println("After")
	})
}
