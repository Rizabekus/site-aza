package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", files))

	mux.HandleFunc("/", Homepage)
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}
}
