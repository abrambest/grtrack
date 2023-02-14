package pkg

import (
	"encoding/json"
	"errors"
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
	Id            int                 `json:"id"`
	Image         string              `json:"image"`
	Name          string              `json:"name"`
	Members       []string            `json:"members"`
	CreationDate  int                 `json:"creationDate"`
	FirstAlbum    string              `json:"firstAlbum"`
	DatesLocation map[string][]string `json:"datesLocations"`
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

func ParsRelation(id string, idNum int) {
	url := "https://groupietrackers.herokuapp.com/api/relation/" + id

	r, err := http.Get(url)
	checkErr(err)

	defer r.Body.Close()

	getRelatoin, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	err = json.Unmarshal(getRelatoin, &Artist[idNum-1])
	checkErr(err)
}

func CheckNum(id int) error {
	if id > Artist[len(Artist)-1].Id || id < 1 {
		return errors.New("Лишнее Братан!")
	}
	return nil
}
