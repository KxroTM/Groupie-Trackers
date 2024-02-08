package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()                     // Create a new application
	w := a.NewWindow("Page d'accueil") // Create a new window

	// Create three buttons
	btn1 := widget.NewButton("Bouton 1", func() {
		// Show a dialog with a message
		dialog.NewInformation("Info", "Vous avez cliqué sur le bouton 1", w).Show()
	})
	btn2 := widget.NewButton("Bouton 2", func() {
		// Show a dialog with a message
		dialog.NewInformation("Info", "Vous avez cliqué sur le bouton 2", w).Show()
	})
	btn3 := widget.NewButton("Bouton 3", func() {
		// Show a dialog with a message
		dialog.NewInformation("Info", "Vous avez cliqué sur le bouton 3", w).Show()
	})

	// Create a vertical box with the buttons
	vBox := container.NewVBox(btn1, btn2, btn3)

	// Create a widget.Card with a title and the buttons
	card := widget.NewCard("Groupie-Trackers", "", vBox)

	// Create a centered container with the widget.Card
	centerContainer := container.NewCenter(card)

	w.SetContent(centerContainer) // Set the window's content
	w.ShowAndRun()                // Show the window and run the application
}
