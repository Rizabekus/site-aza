package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("1")
	mux := http.NewServeMux()
	fmt.Println("12")
	files := http.FileServer(http.Dir("./static"))
	fmt.Println("13")
	mux.Handle("/static/", http.StripPrefix("/static", files))
	fmt.Println("14")
	mux.HandleFunc("/", Homepage)
	mux.HandleFunc("/intro", Intro)
	fmt.Println("15")
	err := http.ListenAndServe(":1515", mux)
	if err != nil {
		fmt.Println("gg")
		log.Fatal(err)
	}
	fmt.Println("??")
	fmt.Println("http://localhost:1515")
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
	http.ServeFile(w, r, "./static/index.html")
}
func Intro(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}
	if r.URL.Path != "/intro" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}
	http.ServeFile(w, r, "./static/home.html")
}
