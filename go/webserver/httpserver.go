package webserver

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world\n")
}
func startHttpServer() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe("127.0.0.1:12345", nil)
}
