package functions

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type API struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

func GetApiUrl() {
	url := "https://groupietrackers.herokuapp.com/api"
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

	var API API

	err = json.Unmarshal(data, &API)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(API)
	fmt.Println(API.Artists)
}
