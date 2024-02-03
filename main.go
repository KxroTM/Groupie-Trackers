package main

import (
	Groupie_Trackers "Groupie_Trackers/tools"
)

func main() {
	// //Appeler la fonction login sur la page de connection
	// Groupie_Trackers.Login("chems", "mdp")
	// //Appeler la fonction register si la fonction login returne false
	// Groupie_Trackers.Register("test", "mdp", "mauvaismdp")
	Groupie_Trackers.AddToFavorites("test", "artiste2")
	Groupie_Trackers.DeleteFavorite("test", "artiste45")
	// Groupie_Trackers.FinalResearch("acdc")
	// for _, data := range datas {
	// 	fmt.Println(data.Name)
	// }
}

func init() {
	Groupie_Trackers.LoadDb()
}
