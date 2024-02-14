package main

import (
	app "Groupie_Trackers/go/app"
	functions "Groupie_Trackers/go/functions"
)

func main() {
	app.Homepage() // Appel de la fonction Homepage du fichier homepage.go
}

func init() {
	functions.LoadDb()
}
