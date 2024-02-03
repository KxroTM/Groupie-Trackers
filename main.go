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
	fmt.Println(len(functions.ArtistbyFirstAlbumDate(ArtistData, "04-10-2008", "04-10-2009")))

	// LocationsData := functions.LocationsData()
	// fmt.Println(LocationsData.Index[0])

	// DatesData := functions.DatesData()
	// fmt.Println(DatesData.Index[0])

	// RelationsData := functions.RelationData()
	// fmt.Println(RelationsData.Index[0])

}
