package main

import (
	"grtrack-mygr/pkg"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Метод запрещен!", 405)
		return
	}
	files := []string{
		"./ui/html/index.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	Info := pkg.Parser()

	err = ts.Execute(w, Info)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internel Server Error", 500)
	}
}

func ShowArtist(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi((path.Base(r.URL.Path)))
	if err != nil {
		log.Println(err)
		return
	}

	if r.URL.Path != "/artist/"+strconv.Itoa(id) {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Метод запрещен!", 405)
		return
	}

	files := []string{
		"./ui/html/artist.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	Info := pkg.Parser()

	err = ts.Execute(w, Info[id-1])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internel Server Error", 500)
	}

	// id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// if err != nil || id < 1 {
	// 	http.NotFound(w, r)
	// 	return
	// }

	// fmt.Fprintf(w, "Отображение артиста с ID %d...", id)
}
