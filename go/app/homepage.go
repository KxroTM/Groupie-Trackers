package app

import (
	"Groupie_Trackers/go/functions"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Homepage() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie-Trackers")

	// Configuration de base pour agrandir les éléments de formulaire
	username := widget.NewEntry()
	username.SetPlaceHolder("Username")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")
	text := canvas.NewText("", color.White)
	text.Alignment = fyne.TextAlignCenter

	loginBtn := widget.NewButton("Login", func() {
		if !functions.Login(username.Text, password.Text) {
			text.Text = "Mauvais mot de passe !"
		} else {
			text.Text = ""
			dialog.ShowInformation("Login", "OUVRE LA PAGE DE L'APPLICATION", myWindow)
			// Open the mainpage.go file
			file, err := os.Open("/go/app/mainpage.go")
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			defer file.Close()
			// Use the file...
		}
	})

	signupBtn := widget.NewButton("Signup", func() {
		// Create signup form elements
		signupUsername := widget.NewEntry()
		signupUsername.SetPlaceHolder("Username")
		signupPassword := widget.NewPasswordEntry()
		signupPassword.SetPlaceHolder("Password")
		signupConfirmPassword := widget.NewPasswordEntry()
		signupConfirmPassword.SetPlaceHolder("Confirm Password")

		signupForm := container.NewVBox(
			signupUsername,
			signupPassword,
			signupConfirmPassword,
		)

		signupDialog := dialog.NewCustom("Signup", "Create Account", signupForm, myWindow)

		signupDialog.Resize(fyne.NewSize(400, 200)) // Agrandissement de la fenêtre

		signupDialog.Show()
	})

	quitBtn := widget.NewButton("Quitter l'application", func() {
		// Ferme l' application
		myApp.Quit()
	})

	loginBtn.Importance = widget.HighImportance  // Augmente l'importance du bouton
	signupBtn.Importance = widget.HighImportance // Augmente l'importance du bouton

	form := container.NewVBox(
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace pour séparer les éléments
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		username,
		password,
		text,
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace pour séparer les éléments
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		container.NewVBox(loginBtn, signupBtn, quitBtn),
	)

	// Ajout d'un titre avec un effet d'ombre
	title := canvas.NewText("Groupie Trackers", theme.ForegroundColor())
	title.TextStyle = fyne.TextStyle{Bold: true} // Texte en gras
	title.TextSize = 42                          // Taille de police plus grande pour le titre

	shadow := canvas.NewText("Groupie Trackers", theme.ShadowColor())
	shadow.TextStyle = fyne.TextStyle{Bold: true} // Texte en gras
	shadow.TextSize = 42
	shadow.Move(fyne.NewPos(2, 2)) // Déplacement léger pour créer l'effet d'ombre

	titleContainer := container.NewWithoutLayout(shadow, title) // Superpose le texte et son ombre

	content := container.NewVBox(
		titleContainer,
		form,
	)
	centeredContent := container.NewCenter(content)

	myWindow.SetContent(centeredContent)
	myWindow.Resize(fyne.NewSize(600, 400)) // Agrandissement de la fenêtre

	myWindow.SetMaster() // Définit la fenêtre principale

	myWindow.ShowAndRun() // Affiche la fenêtre et lance l'application
}
