package main

import (
	"Groupie_Trackers/go/app"
	functions "Groupie_Trackers/go/functions"
)

func main() {
	// start the app by calling the login page
	app.LoginPage(app.MyApp)
	app.MyApp.Run()
}

// Load the database
func init() {
	functions.LoadDb()
}
