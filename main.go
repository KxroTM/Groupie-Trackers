package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("groupie-trackers")

	// Titre
	title := widget.NewLabel("Welcome to My App")

	// Bouttons
	button1 := widget.NewButton("Go to tracker", func() {
		// Button click event handler
		fmt.Println("Button 1 clicked!")
	})

	button2 := widget.NewButton("Click me too!", func() {
		// Button click event handler
		fmt.Println("Button 2 clicked!")
	})

	// Create other widgets for your home page content

	// Create a container to hold all the widgets
	content := container.NewHBox(
		title,
		button1 nil,
		button2 nil,
		// Add other widgets here
	)

	// Set the content of the window to the container
	myWindow.SetContent(content)

	// Show the window and run the app
	myWindow.ShowAndRun()
}
