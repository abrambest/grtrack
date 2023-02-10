package pkg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

type StructArtist struct {
	Id           int    `json:"id"`
	Image        string `json:"image"`
	Name         string //`json:"name"`
	Members      string `json:"members"`
	CreationDate int    `json:"creationDate"`
	FirstAlbum   string
}

var Artist []StructArtist

func Parser() []StructArtist {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	r, err := http.Get(url)
	checkErr(err)

	defer r.Body.Close()

	getContent, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	err = json.Unmarshal(getContent, &Artist)
	checkErr(err)
	return Artist
}
