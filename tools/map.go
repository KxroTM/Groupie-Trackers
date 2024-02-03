package Groupie_Trackers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const apiKey = ""
const bingMapsAPIKey = ""

func AddressToCoordinates(address string) (float64, float64, error) {
	baseURL := "https://api.opencagedata.com/geocode/v1/json"
	apiURL, err := url.Parse(baseURL)
	if err != nil {
		return 0, 0, err
	}

	parameters := url.Values{}
	parameters.Add("key", apiKey)
	parameters.Add("q", address)
	apiURL.RawQuery = parameters.Encode()

	resp, err := http.Get(apiURL.String())
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, 0, err
	}

	results := result["results"].([]interface{})
	if len(results) == 0 {
		return 0, 0, fmt.Errorf("No results found for the given address")
	}

	geometry := results[0].(map[string]interface{})["geometry"].(map[string]interface{})

	lat := geometry["lat"].(float64)

	lng := geometry["lng"].(float64)

	return lat, lng, nil
}

func GenerateMapImageURL(lat, lon float64) string {
	const bingMapsStaticURL = "https://dev.virtualearth.net/REST/v1/Imagery/Map/Road"
	return fmt.Sprintf("%s/%.6f,%.6f/15?mapSize=600,400&key=%s", bingMapsStaticURL, lat, lon, bingMapsAPIKey)
}

// func main() {
// 	address := "Germany Mainz"
// 	lat, lng, err := Groupie_Trackers.AddressToCoordinates(address)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	imageURL := GenerateMapImageURL(lat, lng, bingMapsAPIKey)
// 	fmt.Println(lat, lng)
// 	fmt.Printf("URL de l'image de la carte : %s\n", imageURL)
// }
