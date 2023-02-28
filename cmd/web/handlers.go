package main

import (
	"grtrack-mygr/pkg"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

type Error struct {
	Error string
	Code  int
}

func ErrorPage(w http.ResponseWriter, error string, code int) {
	files := []string{
		"./ui/html/error.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	lf, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)

	err = lf.ExecuteTemplate(w, "error.html", Error{
		Error: error,
		Code:  code,
	})

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
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
		ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	Info := pkg.Parser()
	DlInfo := pkg.ParsRelation()

	type PageData struct {
		Artist   []pkg.StructArtist
		Relation pkg.Relation
	}
	v := PageData{
		Artist:   Info,
		Relation: DlInfo,
	}

	err = ts.Execute(w, v)
	if err != nil {
		log.Println(err.Error())
		ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func ShowArtist(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi((path.Base(r.URL.Path)))
	if err != nil {
		log.Println(err, "ShowArtist get id")
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if r.URL.Path != "/artist/"+strconv.Itoa(id) {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
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
		ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = pkg.CheckNum(id)
	if err != nil {
		log.Println(err.Error())
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err != nil {

		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	Info := pkg.Parser()
	DlInfo := pkg.ParsRelation()

	type PageData struct {
		Artist   pkg.StructArtist
		Relation map[string][]string
	}

	v := PageData{
		Artist:   Info[id-1],
		Relation: DlInfo.Index[id-1].DatesLocation,
	}

	err = ts.Execute(w, v)

	if err != nil {
		log.Println(err.Error())
		ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
