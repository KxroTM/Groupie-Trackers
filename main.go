package main

import (
	"Groupie_Trackers/go/app"
	functions "Groupie_Trackers/go/functions"
)

func main() {
	app.LoginPage(app.MyApp) // Appel de la fonction Homepage du fichier homepage.go
	app.MyApp.Run()
}

func init() {
	functions.LoadDb()
}
