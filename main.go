package main

import "net/http"

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func laurin(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hallo tzbaii"))
}

func main() {
	http.Handle("/lol", &MyHandler{})
	http.HandleFunc("/cool", laurin)
	http.ListenAndServe(":8080", nil)
}
