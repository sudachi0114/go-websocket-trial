package main

import (
	"net/http"
)

func main() {
	// fmt.Println("Hello, world!")
	http.HandleFunc("/ping", handler)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("pong"))
}
