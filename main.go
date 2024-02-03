package main

import (
	"Groupie-Trackers/functions"
	"fmt"
)

func main() {
	// API := functions.ApiUrl()
	// fmt.Println(API)

	ArtistData := functions.ArtistData()
	// fmt.Println(ArtistData[0])

	// fmt.Println(len(functions.ArtistbyCreationDateCheck(ArtistData, []int{1970, 1995})))
	fmt.Println(len(functions.ArtistbyFirstAlbumDateCheck(ArtistData, []string{"29-06-2013", "05-08-1967"})))

	// LocationsData := functions.LocationsData()
	// fmt.Println(LocationsData.Index[0])

	// DatesData := functions.DatesData()
	// fmt.Println(DatesData.Index[0])

	// RelationsData := functions.RelationData()
	// fmt.Println(RelationsData.Index[0])

}
