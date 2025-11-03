package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

type helloWorldHandler struct{}

func (h *helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello,World.")
}

// TODO: どこかのタイミングでテストコードを追加する
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("access_log", "remoteAddr", r.RemoteAddr, "method", r.Method, "url", r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", loggingMiddleware(&helloWorldHandler{}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
