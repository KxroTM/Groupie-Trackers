package functions

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type APIUrl struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type Artist struct {
	ID           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate float64  `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type AllArtists []Artist

type AllLocations struct {
	Index []Locations `json:"index"`
}

type Locations struct {
	ID        int64    `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type AllDates struct {
	Index []Dates `json:"index"`
}

type Dates struct {
	ID    int64    `json:"id"`
	Dates []string `json:"dates"`
}

type AllRelation struct {
	Index []Relation `json:"index"`
}

type Relation struct {
	ID             int64               `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func GetApiData(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func ApiUrl() APIUrl {
	var r APIUrl
	data := GetApiData("https://groupietrackers.herokuapp.com/api")
	json.Unmarshal(data, &r)
	return r
}

func ArtistData() AllArtists {
	var r AllArtists
	data := GetApiData("https://groupietrackers.herokuapp.com/api/artists")
	json.Unmarshal(data, &r)
	return r
}

func LocationsData() AllLocations {
	data := GetApiData("https://groupietrackers.herokuapp.com/api/locations")
	var r AllLocations
	json.Unmarshal(data, &r)
	return r
}

func DatesData() AllDates {
	data := GetApiData("https://groupietrackers.herokuapp.com/api/dates")
	var r AllDates
	json.Unmarshal(data, &r)
	return r
}

func RelationData() AllRelation {
	data := GetApiData("https://groupietrackers.herokuapp.com/api/relation")
	var r AllRelation
	json.Unmarshal(data, &r)
	return r

}
