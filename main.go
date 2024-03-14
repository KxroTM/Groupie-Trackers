package main

import (
	"Groupie_Trackers/go/app"
	functions "Groupie_Trackers/go/functions"
)

func main() {
	app.Homepage() // Appel de la fonction Homepage du fichier homepage.go
	//app.Mainpage() // Appel de la fonction mainPage du fichier mainpage.go
}

func init() {
	functions.LoadDb()
}
