package main

import (
	"grtrack-mygr/pkg"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func NotFoundHandler(w http.ResponseWriter, error string, code int) {
	files := []string{
		"./ui/html/error.html",
		"./ui/html/base.layout.html",
	}
	html, err := template.ParseFiles(files...)
	if err != nil {
		err = html.Execute(w, http.StatusText(code))
		return
	}
	w.WriteHeader(code)
	err = html.Execute(w, http.StatusText(code))

	if err != nil {
		err = html.Execute(w, http.StatusText(code))
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		NotFoundHandler(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
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
		log.Println("ShowArtist get id")
		NotFoundHandler(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if r.URL.Path != "/artist/"+strconv.Itoa(id) {
		NotFoundHandler(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		NotFoundHandler(w, "Метод запрещен!", 405)
		return
	}

	files := []string{
		"./ui/html/artist.html",
		"./ui/html/base.layout.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		NotFoundHandler(w, "Internal Server Error", 500)
		return
	}

	err = pkg.CheckNum(id)
	if err != nil {

		NotFoundHandler(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	Info := pkg.Parser()

	pkg.ParsRelation(strconv.Itoa(id), id)

	err = ts.Execute(w, Info[id-1])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internel Server Error", 500)
	}
}
