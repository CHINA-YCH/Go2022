package test

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

func TestNet(t *testing.T) {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
