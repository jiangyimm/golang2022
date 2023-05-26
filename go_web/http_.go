package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r: %v\n", r.URL.Path)
	if r.URL.Path == "/" {
		sayHelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
