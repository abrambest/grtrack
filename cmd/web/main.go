package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/artist", ShowArtist)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Откройте страницу в браузере http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

// 5. Обработка URL-запросов в Golang - тут рассказано про id
// 8. Получаем доступ к статическим файлам — CSS и JS
// Особенности обработчика статических файлов
