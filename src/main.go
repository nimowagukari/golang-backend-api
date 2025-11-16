package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"
)

type helloWorldHandler struct{}

func (h *helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello,World.")
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(statusCode int) {
	lrw.statusCode = statusCode
	lrw.ResponseWriter.WriteHeader(statusCode)
}

// TODO: どこかのタイミングでテストコードを追加する
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusNotFound}

		next.ServeHTTP(lrw, r)

		elaspedTime := time.Since(start)
		slog.Info("access_log",
			"remoteAddr", r.RemoteAddr,
			"url", r.URL,
			"method", r.Method,
			"statusCode", lrw.statusCode,
			"elaspedSeconds", elaspedTime.Seconds())
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/hello", loggingMiddleware(&helloWorldHandler{}))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
