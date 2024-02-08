package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie Trackers")

	// Configuration de base pour agrandir les éléments de formulaire
	username := widget.NewEntry()
	username.SetPlaceHolder("Username")
	email := widget.NewEntry()
	email.SetPlaceHolder("Email (only for signup)")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	loginBtn := widget.NewButton("Login", func() {
		dialog.ShowInformation("Login", "Login logic here", myWindow)
	})
	signupBtn := widget.NewButton("Signup", func() {
		dialog.ShowInformation("Signup", "Signup logic here", myWindow)
	})
	quitBtn := widget.NewButton("Quitter l'application", func() {
		// Ferme l' application
		myApp.Quit()
	})

	loginBtn.Importance = widget.HighImportance
	signupBtn.Importance = widget.HighImportance

	form := container.NewVBox(
		username,
		email,
		password,
		container.NewVBox(loginBtn, signupBtn, quitBtn),
	)

	// Ajout d'un titre avec un effet d'ombre
	title := canvas.NewText("Groupie Trackers", theme.ForegroundColor())
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.TextSize = 24 // Taille de police plus grande pour le titre

	shadow := canvas.NewText("Groupie Trackers", theme.ShadowColor())
	shadow.TextStyle = fyne.TextStyle{Bold: true}
	shadow.TextSize = 24
	shadow.Move(fyne.NewPos(2, 2)) // Déplacement léger pour créer l'effet d'ombre

	titleContainer := container.NewWithoutLayout(shadow, title) // Superpose le texte et son ombre

	content := container.NewVBox(
		titleContainer,
		form,
	)
	centeredContent := container.NewCenter(content)

	myWindow.SetContent(centeredContent)
	myWindow.Resize(fyne.NewSize(600, 400)) // Agrandissement de la fenêtre

	myWindow.SetMaster()

	myWindow.ShowAndRun()
}
