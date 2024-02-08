package main

import (
	Groupie_Trackers "Groupie_Trackers/tools"
	"fmt"
)

func main() {
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
	data := Groupie_Trackers.SearchByLocation("paris")
	fmt.Println(data)
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
