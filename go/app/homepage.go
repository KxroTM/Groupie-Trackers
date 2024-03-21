package app

import (
	"Groupie_Trackers/go/functions"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var MyApp = app.New()

func LoginPage(app fyne.App) {
	myWindow := app.NewWindow("Groupie-Trackers")

	// Configuration de base pour agrandir les éléments de formulaire
	username := widget.NewEntry()
	username.SetPlaceHolder("Username")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")
	text := canvas.NewText("", color.White)
	text.Alignment = fyne.TextAlignCenter

	loginBtn := widget.NewButton("Login", func() {

		if !functions.Login(username.Text, password.Text) {
			text.Text = "Mot de passe incorrect ou compte inexistant "
		} else {
			text.Text = ""
			Mainpage(app)
			myWindow.Hide()
		}
	})
	signupBtn := widget.NewButton("Signup", func() {
		SignupPage(app)
		myWindow.Hide()
	})

	quitBtn := widget.NewButton("Quitter l'application", func() {
		// Ferme l' application
		app.Quit()
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

	myWindow.SetOnClosed(func() {
		app.Quit()
	})

	myWindow.SetContent(centeredContent)
	myWindow.Resize(fyne.NewSize(600, 400)) // Agrandissement de la fenêtre
	myWindow.CenterOnScreen()

	myWindow.Show() // Affiche la fenêtre et lance l'application
}

func SignupPage(app fyne.App) {
	myWindow := app.NewWindow("Groupie-Trackers")

	// Configuration de base pour agrandir les éléments de formulaire
	signupUsername := widget.NewEntry()
	signupUsername.SetPlaceHolder("Username")
	signupPassword := widget.NewPasswordEntry()
	signupPassword.SetPlaceHolder("Password")
	signupConfirmPassword := widget.NewPasswordEntry()
	signupConfirmPassword.SetPlaceHolder("Password Verification")
	signupText := canvas.NewText("", color.White)
	signupText.Alignment = fyne.TextAlignCenter

	signupBtn := widget.NewButton("Signup", func() {

		if !functions.Register(signupUsername.Text, signupPassword.Text, signupConfirmPassword.Text) {
			signupText.Text = "Mot de passe incorrect ou utilisateur déjà existant"
		} else {
			dialog.ShowInformation("Login", "Compte crée", myWindow)
		}
	})

	loginBtn := widget.NewButton("Login", func() {
		LoginPage(app)
		myWindow.Hide()
	})

	quitBtn := widget.NewButton("Quitter l'application", func() {
		// Ferme l' application
		app.Quit()
	})

	signupBtn.Importance = widget.HighImportance // Augmente l'importance du bouton
	loginBtn.Importance = widget.HighImportance  // Augmente l'importance du bouton

	form := container.NewVBox(
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace pour séparer les éléments
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		signupUsername,
		signupPassword,
		signupConfirmPassword,
		signupText,
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace pour séparer les éléments
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		container.NewVBox(signupBtn, loginBtn, quitBtn),
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

	myWindow.SetOnClosed(func() {
		app.Quit()
	})
	myWindow.SetContent(centeredContent)
	myWindow.Resize(fyne.NewSize(600, 400)) // Agrandissement de la fenêtre
	myWindow.CenterOnScreen()

	myWindow.Show() // Affiche la fenêtre et lance l'application
}
