package main

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mykey"))

func setSession(w http.ResponseWriter, r *http.Request) {

}
func main() {

}
