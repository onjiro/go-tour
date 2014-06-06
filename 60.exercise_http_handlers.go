package main

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s)
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}

func main() {
	//s := String("I'm frayed knot.")
	s := &Struct{"Hello", ":", "Gophers!"}
	http.ListenAndServe("localhost:4000", s)
}
