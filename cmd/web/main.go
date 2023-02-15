package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	addr := flag.String("addr", ":8080", "Сетевой адрес HTTP")
	flag.Parse()
	// example go run ./cmd/web -addr=":3030"

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/artist/", ShowArtist)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	infoLog.Printf("Запуск сервера на http://localhost%s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}

// 5. Обработка URL-запросов в Golang - тут рассказано про id
// 8. Получаем доступ к статическим файлам — CSS и JS
// Особенности обработчика статических файлов
// 12. Внедрение зависимостей в Golang (Dependency Injection) - Сложная тема
