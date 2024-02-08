package main

import (
	Groupie_Trackers "Groupie_Trackers/functions"
	"fmt"
)

func main() {
	// API := functions.ApiUrl()
	// fmt.Println(API)

	// ArtistData := functions.ArtistData()
	// fmt.Println(ArtistData[0])

	// LocationsData := functions.LocationsData()
	// fmt.Println(LocationsData.Index[0])

	// DatesData := functions.DatesData()
	// fmt.Println(DatesData.Index[0])

	// RelationsData := functions.RelationData()
	// fmt.Println(RelationsData.Index[0])

	// //Appeler la fonction login sur la page de connection
	// Groupie_Trackers.Login("chems", "mdp")
	// //Appeler la fonction register si la fonction login returne false
	// Groupie_Trackers.Register("test", "mdp", "mauvaismdp")
	// Groupie_Trackers.AddToFavorites("test", "artiste2")
	// Groupie_Trackers.DeleteFavorite("test", "artiste45")
	// Groupie_Trackers.FinalResearch("acdc")
	// for _, data := range datas {
	// 	fmt.Println(data.Name)
	// }
	fmt.Println(Groupie_Trackers.SearchByMember("ac"))
	fmt.Println("=================================================================")
	fmt.Println(Groupie_Trackers.SearchByName("ac"))
	// address := "Ynov Nanterre France"
	// lat, lng, err := Groupie_Trackers.AddressToCoordinates(address)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// imageURL := Groupie_Trackers.GenerateMapImageURL(lat, lng)
	// fmt.Println(lat, lng)
	// fmt.Printf("URL de l'image de la carte : %s\n", imageURL)
}

func init() {
	Groupie_Trackers.LoadDb()
}
