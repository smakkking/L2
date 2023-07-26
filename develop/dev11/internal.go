package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func isValidDate(ev *Event) bool {
	value, err := time.Parse("2001-05-18", string(ev.Date))

	if err != nil {
		log.Println(err)
		return false
	}
	ev.Time = value

	return true
}

func jsonResult(respondText string) []byte {
	return []byte(fmt.Sprintf(`{"result": %s}`, respondText))
}

func jsonError(respondText string) []byte {
	return []byte(fmt.Sprintf(`{"error": "%s"}`, respondText))
}

func makeJsonRespond(w http.ResponseWriter, code int, data []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(data)
	if err != nil {
		log.Println(err)
	}
}

func MiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println(
			"method", r.Method,
			"path", r.URL.EscapedPath(),
			"duration", time.Since(start),
		)
		next(w, r)
	}

}
