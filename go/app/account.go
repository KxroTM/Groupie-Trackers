package app

import (
	"Groupie_Trackers/go/functions"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func LoginPage(app fyne.App) {
	myWindow := app.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)
	checkRemember = false

	username := widget.NewEntry()
	username.SetPlaceHolder("Nom d'Utilisateur")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Mot de Passe")
	text := canvas.NewText("", color.White)
	text.Alignment = fyne.TextAlignCenter

	loginBtn := widget.NewButton("Se Connecter", func() {

		if !functions.Login(username.Text, password.Text) {
			text.Text = "Mot de passe incorrect ou compte inexistant "
		} else {
			user = functions.UserBuild(username.Text)
			text.Text = ""
			HomePage(app)
			myWindow.Hide()
		}
	})
	signupBtn := widget.NewButton("S'inscrire", func() {
		SignupPage(app)
		myWindow.Hide()
	})

	quitBtn := widget.NewButton("Fermer l'Application", func() {
		// Ferme l' application
		app.Quit()
	})

	check := widget.NewCheck("Se souvenir de moi", func(value bool) {
		if value {
			checkRemember = true
		} else {
			checkRemember = false
		}
	})

	checkButton := container.NewHBox(layout.NewSpacer(), check, layout.NewSpacer())

	loginBtn.Importance = widget.HighImportance  // Augmente l'importance du bouton
	signupBtn.Importance = widget.HighImportance // Augmente l'importance du bouton

	form := container.NewVBox(
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace pour séparer les éléments
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		username,
		password,
		checkButton,
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
	myWindow := app.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)

	// Configuration de base pour agrandir les éléments de formulaire
	signupUsername := widget.NewEntry()
	signupUsername.SetPlaceHolder("Nom d'Utilisateur")
	signupPassword := widget.NewPasswordEntry()
	signupPassword.SetPlaceHolder("Mot de Passe")
	signupConfirmPassword := widget.NewPasswordEntry()
	signupConfirmPassword.SetPlaceHolder("Verification de Mot de Passe")
	signupText := canvas.NewText("", color.White)
	signupText.Alignment = fyne.TextAlignCenter

	signupBtn := widget.NewButton("S'inscrire", func() {

		if !functions.Register(signupUsername.Text, signupPassword.Text, signupConfirmPassword.Text) {
			signupText.Text = "Mot de passe incorrect ou utilisateur déjà existant"
		} else {
			dialog.ShowInformation("Se Connecter", "Compte crée", myWindow)
			LoginPage(app)
			myWindow.Hide()
		}
	})

	loginBtn := widget.NewButton("Se Connecter", func() {
		LoginPage(app)
		myWindow.Hide()
	})

	quitBtn := widget.NewButton("Fermer l'Application", func() {
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

func AccountPage(myApp fyne.App) {
	isPlaying = false

	myWindow := MyApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)
	user = functions.UserBuild(user.Username)

	navBar := createNavBar(myWindow)
	title := canvas.NewText("Mon Compte", color.White)
	title.TextSize = 30
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	subtitle := canvas.NewText("Nom d'utilisateur : "+user.Username, color.White)
	subtitle.TextSize = 20
	subtitle.Alignment = fyne.TextAlignCenter
	spacer := canvas.NewText("", color.White)

	changePasswordButton := widget.NewButton("Changer de mot de passe", func() {
		ChangePasswordPage(myApp)
		myWindow.Hide()
	})
	changePasswordButton.Importance = widget.HighImportance

	if PasswordChange {
		dialog.ShowInformation("Changement de mot de passe", "Mot de passe changé", myWindow)
		PasswordChange = false
	}

	if PpfChange {
		dialog.ShowInformation("Changement de photo de profil", "Photo de profil changée", myWindow)
		PpfChange = false
	}

	changePpfButton := widget.NewButton("Changer de photo de profil", func() {
		ChangePpf(myApp)
		myWindow.Hide()
	})

	playlistButton := widget.NewButton("Mes playlists", func() {
		PlaylistPage(myApp)
		myWindow.Hide()
	})

	if user.Ppf == "" {
		form := container.NewVBox(
			spacer,
			spacer,
			subtitle,
			spacer,
			changePasswordButton,
			changePpfButton,
			playlistButton,
		)

		content := container.NewVBox(
			title,
			form,
		)

		centeredContent := container.NewCenter(content)
		myWindow.SetContent(container.NewBorder(navBar, nil, nil, nil, BackgroundRect, centeredContent))
	} else {

		ppf := loadImageFromURL(user.Ppf) // JARRIVE PAS A REDIMENTIONNER LIMAGE
		form := container.NewVBox(
			spacer,
			changePasswordButton,
			changePpfButton,
			playlistButton,
		)

		content := container.NewVBox(
			title,
			ppf,
			subtitle,
			form,
		)

		centeredContent := container.NewCenter(content)
		myWindow.SetContent(container.NewBorder(navBar, nil, nil, nil, BackgroundRect, centeredContent))
	}

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}

func ChangePasswordPage(myApp fyne.App) {
	myWindow := MyApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)
	user = functions.UserBuild(user.Username)

	navBar := createNavBar(myWindow)

	title := canvas.NewText("Changer de mot de passe", color.White)
	title.TextSize = 28
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	// spacer := canvas.NewText("", color.Transparent)
	oldPassword := widget.NewPasswordEntry()
	oldPassword.SetPlaceHolder("Ancien Mot de Passe")
	newPassword := widget.NewPasswordEntry()
	newPassword.SetPlaceHolder("Nouveau Mot de Passe")
	confirmPassword := widget.NewPasswordEntry()
	confirmPassword.SetPlaceHolder("Confirmer le Nouveau Mot de Passe")
	text := canvas.NewText("", color.White)
	text.Alignment = fyne.TextAlignCenter

	changePasswordButton := widget.NewButton("Changer de mot de passe", func() {
		if !functions.ChangePassword(user.Username, oldPassword.Text, newPassword.Text, confirmPassword.Text) {
			text.Text = "Mot de passe incorrect ou les nouveaux mots de passe ne correspondent pas"
		} else {
			PasswordChange = true
			AccountPage(myApp)
			myWindow.Hide()
		}
	})

	changePasswordButton.Importance = widget.HighImportance
	form := container.NewVBox(
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace pour séparer les éléments
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		oldPassword,
		newPassword,
		confirmPassword,
		container.NewVBox(layout.NewSpacer()),
		text,
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace pour séparer les éléments
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		container.NewVBox(changePasswordButton),
	)

	content := container.NewVBox(
		title,
		form,
	)
	centeredContent := container.NewCenter(content)

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})
	myWindow.SetContent(container.NewBorder(navBar, nil, nil, nil, BackgroundRect, centeredContent))
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}

func ChangePpf(myApp fyne.App) {
	myWindow := myApp.NewWindow("Image Copier")
	myWindow.SetIcon(Icon)
	user = functions.UserBuild(user.Username)

	navBar := createNavBar(myWindow)

	title := canvas.NewText("Changer d'image de profil", color.White)
	title.TextSize = 28
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	entry := widget.NewEntry()
	entry.SetPlaceHolder("URL de l'image")

	copyButton := widget.NewButton("Sélectionner une image", func() {
		if entry.Text == "" {
			dialog.ShowInformation("Erreur", "Veuillez entrer une URL valide", myWindow)
			return
		}
		functions.AddPpf(user.Username, entry.Text)
		PpfChange = true
		AccountPage(myApp)
		myWindow.Hide()
	})
	copyButton.Importance = widget.HighImportance

	form := container.NewVBox(
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace pour séparer les éléments
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		entry,
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace pour séparer les éléments
		container.NewVBox(layout.NewSpacer()), // Ajout d'un espace supplémentaire
		container.NewVBox(copyButton),
	)

	content := container.NewVBox(
		title,
		form,
	)

	centeredContent := container.NewCenter(content)

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})

	myWindow.SetContent(container.NewBorder(navBar, nil, nil, nil, BackgroundRect, centeredContent))
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}
