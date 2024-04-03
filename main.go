package main

import (
	app "Groupie_Trackers/go/app"
	functions "Groupie_Trackers/go/functions"
)

func main() {
	// start the app by calling the login page
	if functions.UserRemember.Username == "" {
		app.LoginPage(app.MyApp)
	} else {
		app.HomePage(app.MyApp)
	}
	app.MyApp.Run()
}

// Load the database
func init() {
	functions.LoadDb()
}
