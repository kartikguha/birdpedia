package main

import (
	"fmt"
	"net/http"
)

type mainHandler struct {
}

func (h mainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "PostForm", r.PostForm)
	fmt.Fprintln(w, "Host", r.Host)
	fmt.Fprintln(w, "Proto", r.Proto)
	fmt.Fprintln(w, "Method", r.Method)
	fmt.Fprintln(w, "ProtoMajor", r.ProtoMajor)
	fmt.Fprintln(w, "ProtoMinor", r.ProtoMinor)
	fmt.Fprintln(w, "RemoteAddr", r.RemoteAddr)
	fmt.Fprintln(w, "RequestURI", r.RequestURI)
}

func init() {
	var h mainHandler
	http.Handle("/", h)
	http.Handle("/hello", h)
}

func main() {
	http.ListenAndServe("0.0.0.0:8086", nil)
}
