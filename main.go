package main

import "net/http"

func main() {
	println("hello, welcome to Go, let's have some fun")

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("great!"))
}
