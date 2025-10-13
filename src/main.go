package main

import (
	"fmt"
	"log"
	"net/http"
)

type helloWorldHandler struct{}

func (h *helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello,World.")
}

func main() {
	http.Handle("/", &helloWorldHandler{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
