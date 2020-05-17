package main

import (
	"fmt"
	"log"
	"net/http"
)

type logger struct {
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("start")
	l.Inner.ServeHTTP(w, req)
	log.Println("finish")
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello\n")
}

func main() {
	f := http.HandlerFunc(hello)
	l := logger{Inner: f}
	http.ListenAndServe(":8000", &l)
}
