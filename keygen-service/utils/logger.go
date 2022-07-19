package utils

import (
	"log"
	"net/http"
	"time"
)

// Logger logs stuff to the console
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s \t%s \t%s \t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func EntityLogger(entity any) {
	log.Printf("%+v", entity)
}
