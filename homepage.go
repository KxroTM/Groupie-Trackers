// Implement the homepage handler
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func homepage() {
	a := app.New()                     // Create a new application
	w := a.NewWindow("Page d'accueil") // Create a new window

	// Create three buttons
	btn1 := widget.NewButton("Se connecter", func() {
		// Create a new window
		w2 := a.NewWindow("Se connecter")

		// Create a form with an entry and a button
		form := &widget.Form{
			Items: []*widget.FormItem{
				{Text: "Nom d'utilisateur", Widget: widget.NewEntry()},
				{Text: "Mot de passe", Widget: widget.NewPasswordEntry()},
			},
			OnSubmit: func() {
				// Declare the form variable
				var form *widget.Form = nil

				// Get the values from the form
				username := form.Items[0].Widget.(*widget.Entry).Text
				password := form.Items[1].Widget.(*widget.Entry).Text

				// Print the values
				println("Nom d'utilisateur:", username)
				println("Mot de passe:", password)
			},
		}

		// Set the window's content to the form
		w2.SetContent(form)

		w2.Show() // Show the window
	})
	btn2 := widget.NewButton("Créer un compte", func() {
		// Create a new window
		w3 := a.NewWindow("Créer un compte")

		// Create a form with an entry and a button
		form := &widget.Form{
			Items: []*widget.FormItem{
				{Text: "Nom d'utilisateur", Widget: widget.NewEntry()},
				{Text: "Mot de passe", Widget: widget.NewPasswordEntry()},
				{Text: "Confirmer le mot de passe", Widget: widget.NewPasswordEntry()},
			},
			OnSubmit: func() {
				// Declare the form variable
				var form *widget.Form = nil

				// Get the values from the form
				username := form.Items[0].Widget.(*widget.Entry).Text
				password := form.Items[1].Widget.(*widget.Entry).Text
				confirmPassword := form.Items[2].Widget.(*widget.Entry).Text

				// Print the values
				println("Nom d'utilisateur:", username)
				println("Mot de passe:", password)
				println("Confirmer le mot de passe:", confirmPassword)
			},
		}

		// Set the window's content to the form
		w3.SetContent(form)

		w3.Show() // Show the window
	})
	btn3 := widget.NewButton("Quitter l'application", func() {
		// Close the application
		a.Quit()
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
