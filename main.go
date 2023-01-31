package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	log.Println("Откройте страницу в браузере http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

// 2.Основы веб-приложений на Golang
