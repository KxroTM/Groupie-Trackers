package app

import (
	"Groupie_Trackers/go/functions"
	"image/color"
	"strconv"

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
var user *functions.Account
var checkRemember = false
var Icon, _ = fyne.LoadResourceFromPath("./Icon.png")
var BackgroundRect = canvas.NewRectangle(color.RGBA{R: 16, G: 16, B: 16, A: 255}) //Background color (dark grey)
var PasswordChange = false
var PpfChange = false
var isPlaying = false
var playlistCreate = false

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

func SearchPage(myApp fyne.App) {
	myWindow := myApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)
	isPlaying = false

	filterButton := widget.NewButton("Recherche avec Filtre", func() {
		FilterPage(MyApp)
		myWindow.Hide()
	})

	navBar := createNavBar(myWindow)
	artists := functions.ArtistData()

	artistsGrid := createArtistsGrid(artists, myWindow)
	gridContainer := container.NewStack() // Utilisation de NewMax pour pouvoir rafraîchir dynamiquement le contenu
	gridContainer.Add(artistsGrid)

	searchBar := createSearchBar(artists, gridContainer, myWindow)
	topContent := container.NewVBox(navBar, filterButton, searchBar)

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})

	myWindow.SetContent(container.NewBorder(topContent, nil, nil, nil, BackgroundRect, gridContainer))
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}

func Propospage(myApp fyne.App) {
	myWindow := MyApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)

	text := canvas.NewText("Groupie Tracker : Projet de création d'une application pour suivre les groupes de musique.", color.White)
	text.Alignment = fyne.TextAlignCenter

	text2 := canvas.NewText("Utilisation d'une API pour manipuler les données sur les artistes, les lieux et les dates de concerts.", color.White)
	text2.Alignment = fyne.TextAlignCenter

	text3 := canvas.NewText("Objectif : Créer une application conviviale affichant les informations sur les groupes via diverses visualisations de données.", color.White)
	text3.Alignment = fyne.TextAlignCenter

	text4 := canvas.NewText("Focus sur la création d'événements/actions interactifs, comme les appels client-serveur.", color.White)
	text4.Alignment = fyne.TextAlignCenter

	text5 := canvas.NewText("Implémentation d'une géolocalisation des concerts sur une carte.", color.White)
	text5.Alignment = fyne.TextAlignCenter

	text6 := canvas.NewText("Création d'une barre de recherche permettant de rechercher des membres ou des artistes.", color.White)
	text6.Alignment = fyne.TextAlignCenter

	text7 := canvas.NewText("Intégration de filtres pour afficher les groupes selon différents critères.", color.White)
	text7.Alignment = fyne.TextAlignCenter

	text8 := canvas.NewText("Respect des 8 règles d'interface de Schneiderman pour les visualisations de données.", color.White)
	text8.Alignment = fyne.TextAlignCenter

	text9 := canvas.NewText("Développement en Go, en suivant les bonnes pratiques et en utilisant uniquement les packages standard et Fyne.", color.White)
	text9.Alignment = fyne.TextAlignCenter

	text10 := canvas.NewText("Aides et ressources fournies pour la mise en œuvre, y compris des exemples d'API RESTful et de GUI avec Fyne.", color.White)
	text10.Alignment = fyne.TextAlignCenter

	text11 := canvas.NewText("Bonus possibles : Intégration d'un lecteur Spotify et création d'un système de favoris.", color.White)
	text11.Alignment = fyne.TextAlignCenter

	navBar := createNavBar(myWindow)
	content := container.NewStack(container.NewBorder(navBar, nil, nil, nil, container.NewCenter(container.NewVBox(text, text2, text3, text4, text5, text6, text7, text8, text9, text10, text11))))

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})

	myWindow.SetContent(content)
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}

func Contactpage(myApp fyne.App) {
	myWindow := MyApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)

	text := canvas.NewText("Contactez-nous à l'adresse suivante:", color.White)
	text.Alignment = fyne.TextAlignCenter

	email := canvas.NewText("Email: GroupieTrackers@Ynov.com", color.White)
	email.Alignment = fyne.TextAlignCenter

	navBar := createNavBar(myWindow)

	content := container.NewStack(container.NewBorder(navBar, nil, nil, nil, container.NewCenter(container.NewVBox(text, email))))

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})
	myWindow.SetContent(content)
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}

func ArtistPage(artist functions.Artist, myApp fyne.App) {
	functions.AddHistory(user.Username, artist.Name)
	myWindow := MyApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)
	user = functions.UserBuild(user.Username)
	var content *fyne.Container
	var playButtonText string
	var favoriteButton *fyne.Container
	var selecter *widget.Select

	navBar := createNavBar(myWindow)

	image := loadImageFromURL(artist.Image)
	image.FillMode = canvas.ImageFillContain

	name := canvas.NewText("Nom : "+artist.Name, color.White)
	members := ""
	members2 := ""

	if len(artist.Members) > 4 {
		for i := 0; i < len(artist.Members)/2; i++ {
			if i == 0 {
				members += artist.Members[i]
			} else {
				members += ", " + artist.Members[i]
			}
		}

		for i := len(artist.Members) / 2; i < len(artist.Members); i++ {
			if i == len(artist.Members) {
				members2 += artist.Members[i]
			} else if i == len(artist.Members)/2 {
				members2 += artist.Members[i]
			} else {
				members2 += ", " + artist.Members[i]
			}
		}
	} else {
		for i := 0; i < len(artist.Members); i++ {
			if i == 0 {
				members += artist.Members[i]
			} else if i == len(artist.Members) {
				members += artist.Members[i]
			} else {
				members += ", " + artist.Members[i]
			}
		}
	}

	member := canvas.NewText("Membres : "+members, color.White)
	member2 := canvas.NewText(members2, color.White)

	creationDate := canvas.NewText("Date de Création : "+strconv.Itoa(int(artist.CreationDate)), color.White)
	album := canvas.NewText("Premier Album : "+artist.FirstAlbum, color.White)
	concert := widget.NewButton("Concerts", func() {
		ConcertPage(artist, MyApp)
		myWindow.Hide()
	})
	concertButton := container.NewHBox(layout.NewSpacer(), concert, layout.NewSpacer())
	spacer := canvas.NewText("              ", color.White)

	if !isPlaying {
		playButtonText = "Écouter un extrait"
	} else {
		playButtonText = "Arrêter l'extrait"
	}

	play := widget.NewButton(playButtonText, func() {
		if !isPlaying {
			isPlaying = true
		} else {
			isPlaying = false
		}
		ArtistPage(artist, myApp)
		myWindow.Hide()
	})

	playlistbutton := widget.NewButton("Ajouter à une playlist", func() {
		CreatePlaylistPage(myApp, artist)
		myWindow.Hide()
	})

	var playlistList []string
	playlistList = append(playlistList, "Créer à une playlist")
	for _, playlist := range user.Playlists.Playlist {
		playlistList = append(playlistList, playlist.Name)
	}

	selecter = widget.NewSelect(playlistList, func(s string) {
		if s == "Créer à une playlist" {
			CreatePlaylistPage(myApp, artist)
			myWindow.Hide()
		} else {
			if functions.IsInPlaylist(user.Username, s, artist.Name) {
				dialog.ShowInformation("Erreur", "Déjà dans la playlist "+s, myWindow)
				selecter.Selected = "Ajouter à une playlist"
				selecter.Refresh()
				return
			} else {
				dialog.ShowInformation("Ajout à la playlist", "Ajouté à la playlist "+s, myWindow)
				functions.AddSongToPlaylist(user.Username, s, artist.Name)
				selecter.Selected = "Ajouter à une playlist"
				selecter.Refresh()
			}
		}
	})

	selecter.PlaceHolder = "Ajouter à une playlist"

	playButton := container.NewHBox(layout.NewSpacer(), play, layout.NewSpacer())

	name.Alignment = fyne.TextAlignCenter
	member.Alignment = fyne.TextAlignCenter
	member2.Alignment = fyne.TextAlignCenter
	creationDate.Alignment = fyne.TextAlignCenter
	album.Alignment = fyne.TextAlignCenter
	txt := canvas.NewText("", color.White)
	embed := SpotifyEmbed(artist)

	if playlistCreate {
		dialog.ShowInformation("Création de playlist", "Playlist créée", myWindow)
		playlistCreate = false
	}

	if len(artist.Members) > 4 {
		if !functions.IsInFavorite(user.Username, artist.Name) {
			dislike := createCustomIcon("src/dislike.png")
			favorite := widget.NewButtonWithIcon("", dislike, func() {
				functions.AddToFavorites(user.Username, artist.Name)
				ArtistPage(artist, myApp)
				myWindow.Hide()
			})
			if len(user.Playlists.Playlist) == 0 {
				favoriteButton = container.NewHBox(layout.NewSpacer(), layout.NewSpacer(), playlistbutton, favorite, spacer)
			} else {
				favoriteButton = container.NewHBox(layout.NewSpacer(), layout.NewSpacer(), selecter, favorite, spacer)
			}
			imgContentCenter := container.NewCenter(container.NewHBox(layout.NewSpacer(), image, layout.NewSpacer()))
			form := container.NewVScroll(container.NewVBox(txt, txt, favoriteButton, txt, txt, imgContentCenter, txt, playButton, txt, name, member, member2, creationDate, album, txt, concertButton, txt))
			if !isPlaying {
				content = container.NewBorder(navBar, nil, nil, nil, BackgroundRect, form)
			} else {
				content = container.NewBorder(navBar, embed, nil, nil, BackgroundRect, form)
			}
		} else {
			like := createCustomIcon("src/like.png")
			favorite := widget.NewButtonWithIcon("", like, func() {
				functions.DeleteFavorite(user.Username, artist.Name)
				ArtistPage(artist, myApp)
				myWindow.Hide()
			})
			if len(user.Playlists.Playlist) == 0 {
				favoriteButton = container.NewHBox(layout.NewSpacer(), layout.NewSpacer(), playlistbutton, favorite, spacer)
			} else {
				favoriteButton = container.NewHBox(layout.NewSpacer(), layout.NewSpacer(), selecter, favorite, spacer)
			}
			imgContentCenter := container.NewCenter(container.NewHBox(layout.NewSpacer(), image, layout.NewSpacer()))
			form := container.NewVScroll(container.NewVBox(txt, txt, favoriteButton, txt, txt, imgContentCenter, txt, playButton, txt, name, member, member2, creationDate, album, txt, concertButton, txt))
			if !isPlaying {
				content = container.NewBorder(navBar, nil, nil, nil, BackgroundRect, form)
			} else {
				content = container.NewBorder(navBar, embed, nil, nil, BackgroundRect, form)
			}
		}
	} else {
		if !functions.IsInFavorite(user.Username, artist.Name) {
			dislike := createCustomIcon("src/dislike.png")
			favorite := widget.NewButtonWithIcon("", dislike, func() {
				functions.AddToFavorites(user.Username, artist.Name)
				ArtistPage(artist, myApp)
				myWindow.Hide()
			})
			if len(user.Playlists.Playlist) == 0 {
				favoriteButton = container.NewHBox(layout.NewSpacer(), layout.NewSpacer(), playlistbutton, favorite, spacer)
			} else {
				favoriteButton = container.NewHBox(layout.NewSpacer(), layout.NewSpacer(), selecter, favorite, spacer)
			}
			imgContentCenter := container.NewCenter(container.NewHBox(layout.NewSpacer(), image, layout.NewSpacer()))
			form := container.NewVScroll(container.NewVBox(txt, txt, favoriteButton, txt, txt, imgContentCenter, txt, playButton, txt, name, member, creationDate, album, txt, concertButton, txt))
			if !isPlaying {
				content = container.NewBorder(navBar, nil, nil, nil, BackgroundRect, form)
			} else {
				content = container.NewBorder(navBar, embed, nil, nil, BackgroundRect, form)
			}
		} else {
			like := createCustomIcon("src/like.png")
			favorite := widget.NewButtonWithIcon("", like, func() {
				functions.DeleteFavorite(user.Username, artist.Name)
				ArtistPage(artist, myApp)
				myWindow.Hide()
			})
			if len(user.Playlists.Playlist) == 0 {
				favoriteButton = container.NewHBox(layout.NewSpacer(), layout.NewSpacer(), playlistbutton, favorite, spacer)
			} else {
				favoriteButton = container.NewHBox(layout.NewSpacer(), layout.NewSpacer(), selecter, favorite, spacer)
			}
			imgContentCenter := container.NewCenter(container.NewHBox(layout.NewSpacer(), image, layout.NewSpacer()))
			form := container.NewVScroll(container.NewVBox(txt, txt, favoriteButton, txt, txt, imgContentCenter, txt, playButton, txt, name, member, creationDate, album, txt, concertButton, txt))
			if !isPlaying {
				content = container.NewBorder(navBar, nil, nil, nil, BackgroundRect, form)
			} else {
				content = container.NewBorder(navBar, embed, nil, nil, BackgroundRect, form)
			}
		}
	}

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})

	myWindow.SetContent(content)
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}

func HomePage(myApp fyne.App) {
	isPlaying = false

	if functions.UserRemember.Username == "" {
		user = functions.UserBuild(user.Username)
	} else {
		user = functions.UserBuild(functions.UserRemember.Username)
	}

	myWindow := MyApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)

	if checkRemember {
		functions.RememberMe(user.Username)
	}

	navBar := createNavBar(myWindow)
	rdmBar := createRandomArtistsGrid(myWindow)
	lastAlbumBar := createCustomArtistsGrid(myWindow, functions.SortByFirstAlbumDescending(functions.ArtistData()))
	firstAlbumBar := createCustomArtistsGrid(myWindow, functions.SortByFirstAlbumAscending(functions.ArtistData()))
	favoriteBar := createFavoriteGrid(myWindow, *user)
	historyBar := createHistoryGrid(myWindow, *user)
	spacer := canvas.NewText("  ", color.White)
	subtitle := canvas.NewText("  Découvrez de nouveaux artistes", color.White)
	subtitle.TextSize = 22
	subtitle.TextStyle = fyne.TextStyle{Bold: true}
	subtitle2 := canvas.NewText("  Découvrez les dernières sorties", color.White)
	subtitle2.TextSize = 22
	subtitle2.TextStyle = fyne.TextStyle{Bold: true}
	subtitle3 := canvas.NewText("  Redécouvrez le meilleur des vieux albums", color.White)
	subtitle3.TextSize = 22
	subtitle3.TextStyle = fyne.TextStyle{Bold: true}
	subtitle4 := canvas.NewText("  Vos favoris", color.White)
	subtitle4.TextSize = 22
	subtitle4.TextStyle = fyne.TextStyle{Bold: true}
	subtitle5 := canvas.NewText("  Récemment consulté", color.White)
	subtitle5.TextSize = 22
	subtitle5.TextStyle = fyne.TextStyle{Bold: true}

	favorite := container.NewHBox(subtitle4, layout.NewSpacer())

	content := container.NewVBox(spacer, spacer, spacer,
		subtitle, spacer, rdmBar, spacer, favorite, spacer, favoriteBar, spacer,
		subtitle2, spacer, lastAlbumBar, spacer, subtitle3, spacer, firstAlbumBar,
		spacer, subtitle5, spacer, historyBar, spacer)

	scrollContainer := container.NewVScroll(content)

	myWindow.SetContent(container.NewBorder(navBar, nil, nil, nil, BackgroundRect, scrollContainer))

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}

func FavoritePage(myApp fyne.App) {
	isPlaying = false

	myWindow := MyApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)
	user = functions.UserBuild(user.Username)

	spacer := canvas.NewText("", color.White)
	navBar := createNavBar(myWindow)
	favoriteBar := createAllFavoriteGrid(myWindow, *user)
	title := canvas.NewText("Mes favoris", color.White)
	title.TextSize = 30
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	content := container.NewVBox(spacer, spacer, title, spacer, favoriteBar)
	scrollContainer := container.NewVScroll(content)

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})
	myWindow.SetContent(container.NewBorder(navBar, nil, nil, nil, BackgroundRect, scrollContainer))
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
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
	myWindow.SetContent(container.NewBorder(navBar, nil, nil, nil, centeredContent))
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

	myWindow.SetContent(container.NewBorder(navBar, nil, nil, nil, centeredContent))
	myWindow.SetContent(form)
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.Show()
}

func PlaylistPage(myApp fyne.App) {
	myWindow := myApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)
	user = functions.UserBuild(user.Username)
	var playlistGrid fyne.CanvasObject

	navBar := createNavBar(myWindow)

	title := canvas.NewText("Mes Playlists", color.White)
	spacer := canvas.NewText("   ", color.White)
	title.TextSize = 30
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	playlistsContainer := container.NewVBox()

	for _, playlist := range user.Playlists.Playlist {
		Title := canvas.NewText(" "+playlist.Name, color.White)
		Title.TextSize = 22
		Title.TextStyle = fyne.TextStyle{Bold: true}
		playlistButton := widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
			OnePlaylistPage(myApp, *user, playlist.Name)
			myWindow.Hide()
		})
		deleteButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
			functions.DeletePlaylist(user.Username, playlist.Name)
			PlaylistPage(myApp)
			myWindow.Hide()
		})

		titleContent := container.NewHBox(Title, layout.NewSpacer(), deleteButton, playlistButton, spacer)
		playlistGrid = createPlaylistGrid(myWindow, *user, playlist.Name)
		playlistsContainer.Add(container.NewVBox(titleContent, spacer, playlistGrid, spacer))
	}

	fullContent := container.NewVScroll(container.NewVBox(spacer, title, spacer, playlistsContainer, spacer))
	content := container.NewStack(container.NewBorder(navBar, nil, nil, nil, BackgroundRect, fullContent))

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})
	myWindow.SetContent(content)
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}

func OnePlaylistPage(myApp fyne.App, user functions.Account, playlist string) {
	myWindow := myApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)

	navBar := createNavBar(myWindow)

	playlistGrid := createFullPlaylistGrid(myWindow, user, playlist)

	content := container.NewStack(container.NewBorder(navBar, nil, nil, nil, BackgroundRect, playlistGrid))

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})
	myWindow.SetContent(content)
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}

func CreatePlaylistPage(myApp fyne.App, artist functions.Artist) {
	myWindow := myApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)

	navBar := createNavBar(myWindow)

	playlistName := widget.NewEntry()
	playlistName.SetPlaceHolder("Nom de la playlist")
	spacer := canvas.NewText("==============================", color.Transparent)
	text := canvas.NewText("", color.White)
	text.Alignment = fyne.TextAlignCenter

	submitButton := widget.NewButton("Créer", func() {
		for _, playlist := range user.Playlists.Playlist {
			if playlist.Name == playlistName.Text {
				text.Text = "Nom de playlist déjà existant"
				return
			}
		}
		functions.CreatePlaylist(user.Username, playlistName.Text)
		playlistCreate = true
		ArtistPage(artist, myApp)
		myWindow.Hide()
	})

	submitButton.Importance = widget.HighImportance

	form := container.NewVBox(
		container.NewVBox(layout.NewSpacer()),
		container.NewVBox(layout.NewSpacer()),
		container.NewVBox(layout.NewSpacer()),
		container.NewVBox(layout.NewSpacer()),
		spacer,
		playlistName,
		text,
		container.NewVBox(layout.NewSpacer()),
		container.NewVBox(layout.NewSpacer()),
		container.NewVBox(submitButton),
	)

	centeredContent := container.NewCenter(form)

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})

	myWindow.SetContent(container.NewBorder(navBar, nil, nil, nil, BackgroundRect, centeredContent))
	myWindow.Resize(fyne.NewSize(600, 400))
	myWindow.CenterOnScreen()

	myWindow.Show()
}
