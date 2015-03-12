package main

import (
	"net/http"
)

func initHTTP() error {
	http.HandleFunc("*", httpIndex)
	return nil
}

func httpIndex(w http.ResponseWriter, r *http.Request) {

}
