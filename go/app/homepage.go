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
var Icon, _ = fyne.LoadResourceFromPath("./Icon.png")

func LoginPage(app fyne.App) {
	myWindow := app.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)
	// Configuration de base pour agrandir les éléments de formulaire
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

	myWindow.SetContent(container.NewBorder(topContent, nil, nil, nil, gridContainer))
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}

func FilterPage(myApp fyne.App) {
	myWindow := myApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)

	navBar := createNavBar(myWindow)
	researchButton := widget.NewButton("Recherche sans Filtre", func() {
		FilterPage(MyApp)
		myWindow.Hide()
	})

	// Fonction pour mettre à jour le Label avec l'année correspondant à la valeur du slider
	updateLabelYear := func(label *widget.Label, value float64) {
		year := int(value)
		label.SetText(strconv.Itoa(year))
	}

	// Créer des Labels pour afficher les années
	labelCreationDateStart := widget.NewLabel("1958")
	labelCreationDateEnd := widget.NewLabel("2015")

	sliderCreationDateStart := widget.NewSlider(1958, 2015) // 1958 for Bee Gees and 2015 for Juice Wrld
	sliderCreationDateEnd := widget.NewSlider(1958, 2015)   // 1958 for Bee Gees and 2015 for Juice Wrld

	sliderCreationDateStart.SetValue(1958)
	sliderCreationDateEnd.SetValue(2015)

	// Mettre à jour les Labels à chaque fois que la valeur des sliders change
	sliderCreationDateStart.OnChanged = func(value float64) {
		updateLabelYear(labelCreationDateStart, value)
	}
	sliderCreationDateEnd.OnChanged = func(value float64) {
		updateLabelYear(labelCreationDateEnd, value)
	}

	creationDateRange := container.NewVBox(
		widget.NewLabel("Date de Création :"),
		container.NewHBox(
			sliderCreationDateStart,
			labelCreationDateStart,
		),
		container.NewHBox(
			sliderCreationDateEnd,
			labelCreationDateEnd,
		),
	)

	// Déclaration et initialisation des sliders pour First Album Date Range

	labelFirstAlbumStart := widget.NewLabel("1967")
	labelFirstAlbumDateEnd := widget.NewLabel("2018")

	sliderFirstAlbumStart := widget.NewSlider(1967, 2018) // 1967 for Pink Floyd and 2018 for Juice Wrld
	sliderFirstAlbumEnd := widget.NewSlider(1967, 2018)   // 1967 for Pink Floyd and 2018 for Juice Wrld

	sliderFirstAlbumStart.SetValue(1967)
	sliderFirstAlbumEnd.SetValue(2018)

	// Mettre à jour les Labels à chaque fois que la valeur des sliders change
	sliderFirstAlbumStart.OnChanged = func(value float64) {
		updateLabelYear(labelFirstAlbumStart, value)
	}
	sliderFirstAlbumEnd.OnChanged = func(value float64) {
		updateLabelYear(labelFirstAlbumDateEnd, value)
	}

	firstAlbumRange := container.NewVBox(
		widget.NewLabel("Publication du Premier Album :"),
		container.NewHBox(
			sliderFirstAlbumStart,
			labelFirstAlbumStart,
		),
		container.NewHBox(
			sliderFirstAlbumEnd,
			labelFirstAlbumDateEnd,
		),
	)

	// Déclaration et initialisation de l'entry pour Number of Members
	oneM := widget.NewCheck("1", func(checked bool) {})
	twoM := widget.NewCheck("2", func(checked bool) {})
	threeM := widget.NewCheck("3", func(checked bool) {})
	fourM := widget.NewCheck("4", func(checked bool) {})
	fiveM := widget.NewCheck("5", func(checked bool) {})
	sixM := widget.NewCheck("6", func(checked bool) {})
	sevenM := widget.NewCheck("7", func(checked bool) {})

	numMembers := container.NewVBox(
		widget.NewLabel("Nombre de Membres :"),
		container.NewHBox(
			container.NewHBox(
				oneM,
				twoM,
				threeM,
				fourM,
			),
			container.NewHBox(
				fiveM,
				sixM,
				sevenM,
			),
		),
	)

	DateRange := container.NewHBox(
		creationDateRange,
		firstAlbumRange,
		numMembers,
	)

	// Déclaration et initialisation des checkboxes pour Locations
	checkUSA := widget.NewCheck("États-Unis", func(checked bool) {})
	checkJapan := widget.NewCheck("Japon", func(checked bool) {})
	checkNewZealand := widget.NewCheck("Nouvelle-Zélande", func(checked bool) {})
	checkMexico := widget.NewCheck("Mexique", func(checked bool) {})
	checkFrenchPolynesia := widget.NewCheck("Polynésie française", func(checked bool) {})
	checkNewCaledonia := widget.NewCheck("Nouvelle-Calédonie", func(checked bool) {})
	checkUK := widget.NewCheck("Royaume-Uni", func(checked bool) {})
	checkSwitzerland := widget.NewCheck("Suisse", func(checked bool) {})
	checkFrance := widget.NewCheck("France", func(checked bool) {})
	checkAustralia := widget.NewCheck("Australie", func(checked bool) {})
	checkIndonesia := widget.NewCheck("Indonésie", func(checked bool) {})
	checkSlovakia := widget.NewCheck("Slovaquie", func(checked bool) {})
	checkHungary := widget.NewCheck("Hongrie", func(checked bool) {})
	checkBelarus := widget.NewCheck("Biélorussie", func(checked bool) {})
	checkBrazil := widget.NewCheck("Brésil", func(checked bool) {})
	checkColombia := widget.NewCheck("Colombie", func(checked bool) {})
	checkDenmark := widget.NewCheck("Danemark", func(checked bool) {})
	checkGermany := widget.NewCheck("Allemagne", func(checked bool) {})
	checkQatar := widget.NewCheck("Qatar", func(checked bool) {})
	checkIndia := widget.NewCheck("Inde", func(checked bool) {})
	checkUnitedArabEmirates := widget.NewCheck("Émirats arabes unis", func(checked bool) {})
	checkCanada := widget.NewCheck("Canada", func(checked bool) {})
	checkFinland := widget.NewCheck("Finlande", func(checked bool) {})
	checkChile := widget.NewCheck("Chili", func(checked bool) {})
	checkArgentina := widget.NewCheck("Argentine", func(checked bool) {})
	checkPeru := widget.NewCheck("Pérou", func(checked bool) {})
	checkPortugal := widget.NewCheck("Portugal", func(checked bool) {})
	checkSweden := widget.NewCheck("Suède", func(checked bool) {})
	checkNorway := widget.NewCheck("Norvège", func(checked bool) {})
	checkNetherlands := widget.NewCheck("Pays-Bas", func(checked bool) {})
	checkChina := widget.NewCheck("Chine", func(checked bool) {})
	checkPhilippines := widget.NewCheck("Philippines", func(checked bool) {})
	checkSaudiArabia := widget.NewCheck("Arabie saoudite", func(checked bool) {})
	checkTaiwan := widget.NewCheck("Taïwan", func(checked bool) {})
	checkSouthKorea := widget.NewCheck("Corée du Sud", func(checked bool) {})
	checkCostaRica := widget.NewCheck("Costa Rica", func(checked bool) {})
	checkAustria := widget.NewCheck("Autriche", func(checked bool) {})

	locations := container.NewVBox(
		widget.NewLabel("Localisation des concerts :"),
		container.NewHBox(
			checkUSA,
			checkJapan,
			checkNewZealand,
			checkMexico,
			checkFrenchPolynesia,
			checkNewCaledonia,
		),
		container.NewHBox(
			checkUK,
			checkSwitzerland,
			checkFrance,
			checkAustralia,
			checkIndonesia,
			checkSlovakia,
		),
		container.NewHBox(
			checkHungary,
			checkBelarus,
			checkBrazil,
			checkColombia,
			checkDenmark,
			checkGermany,
		),
		container.NewHBox(
			checkQatar,
			checkIndia,
			checkUnitedArabEmirates,
			checkCanada,
			checkFinland,
			checkChile,
		),
		container.NewHBox(
			checkArgentina,
			checkPeru,
			checkPortugal,
			checkSweden,
			checkNorway,
			checkNetherlands,
		),
		container.NewHBox(
			checkChina,
			checkPhilippines,
			checkSaudiArabia,
			checkTaiwan,
			checkSouthKorea,
			checkCostaRica,
			checkAustria,
		),
	)

	applyButton := widget.NewButton("Apply Filters", func() {

	})

	resetButton := widget.NewButton("Reset Filters", func() {
		// Ici, vous pouvez maintenant accéder directement aux variables
		sliderCreationDateStart.SetValue(1958)
		sliderCreationDateEnd.SetValue(2015)
		sliderFirstAlbumStart.SetValue(1967)
		sliderFirstAlbumEnd.SetValue(2018)

		oneM.SetChecked(false)
		twoM.SetChecked(false)
		threeM.SetChecked(false)
		fourM.SetChecked(false)
		fiveM.SetChecked(false)
		sixM.SetChecked(false)
		sevenM.SetChecked(false)

		// checkUSA.SetChecked(false)
		// checkUK.SetChecked(false)
		// checkFR.SetChecked(false)
	})

	filterContainer := container.NewVBox(
		DateRange,
		locations,
		container.NewHBox(applyButton, resetButton),
	)

	artists := functions.ArtistData()

	artistsGrid := createArtistsGrid(artists, myWindow)
	gridContainer := container.NewVBox(artistsGrid) // Ajouter artistsGrid à gridContainer

	scrollContainer := container.NewVScroll(gridContainer) // Placer gridContainer dans un conteneur défilable

	topContent := container.NewVBox(navBar, researchButton, filterContainer)

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})

	myWindow.SetContent(container.NewBorder(topContent, nil, nil, nil, scrollContainer)) // Utiliser scrollContainer à la place de gridContainer
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
	myWindow := MyApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)
	user = functions.UserBuild(user.Username)

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

	name.Alignment = fyne.TextAlignCenter
	member.Alignment = fyne.TextAlignCenter
	member2.Alignment = fyne.TextAlignCenter
	creationDate.Alignment = fyne.TextAlignCenter
	album.Alignment = fyne.TextAlignCenter
	txt := canvas.NewText("", color.White)

	if len(artist.Members) > 4 {
		if !functions.IsInFavorite(user.Username, artist.Name) {
			favorite := widget.NewButton("Ajouter aux favoris", func() {
				functions.AddToFavorites(user.Username, artist.Name)
				ArtistPage(artist, myApp)
				myWindow.Hide()
			})

			favoriteButton := container.NewHBox(layout.NewSpacer(), favorite, layout.NewSpacer())
			form := container.NewVBox(navBar, txt, txt, txt, txt, image, favoriteButton, txt, name, member, member2, creationDate, album, txt, concertButton)
			myWindow.SetContent(form)
		} else {
			favorite := widget.NewButton("Retirer des favoris", func() {
				functions.DeleteFavorite(user.Username, artist.Name)
				ArtistPage(artist, myApp)
				myWindow.Hide()
			})

			favoriteButton := container.NewHBox(layout.NewSpacer(), favorite, layout.NewSpacer())
			form := container.NewVBox(navBar, txt, txt, txt, txt, image, favoriteButton, txt, name, member, member2, creationDate, album, txt, concertButton)
			myWindow.SetContent(form)

		}
	} else {
		if !functions.IsInFavorite(user.Username, artist.Name) {
			favorite := widget.NewButton("Ajouter aux favoris", func() {
				functions.AddToFavorites(user.Username, artist.Name)
				ArtistPage(artist, myApp)
				myWindow.Hide()
			})

			favoriteButton := container.NewHBox(layout.NewSpacer(), favorite, layout.NewSpacer())
			form := container.NewVBox(navBar, txt, txt, txt, txt, image, favoriteButton, txt, name, member, creationDate, album, txt, concertButton)
			myWindow.SetContent(form)
		} else {
			favorite := widget.NewButton("Retirer des favoris", func() {
				functions.DeleteFavorite(user.Username, artist.Name)
				ArtistPage(artist, myApp)
				myWindow.Hide()
			})

			favoriteButton := container.NewHBox(layout.NewSpacer(), favorite, layout.NewSpacer())
			form := container.NewVBox(navBar, txt, txt, txt, txt, image, favoriteButton, txt, name, member, creationDate, album, txt, concertButton)
			myWindow.SetContent(form)

		}
	}

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}

func ConcertPage(artist functions.Artist, myApp fyne.App) {
	myWindow := MyApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)

	navBar := createNavBar(myWindow)

	content := container.NewStack(container.NewBorder(navBar, nil, nil, nil))

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})
	myWindow.SetContent(content)
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}

func HomePage(myApp fyne.App) {
	user = functions.UserBuild(user.Username)
	myWindow := MyApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)

	navBar := createNavBar(myWindow)
	rdmBar := createRandomArtistsGrid(myWindow)
	lastAlbumBar := createCustomArtistsGrid(myWindow, functions.SortByFirstAlbumDescending(functions.ArtistData()))
	firstAlbumBar := createCustomArtistsGrid(myWindow, functions.SortByFirstAlbumAscending(functions.ArtistData()))
	favoriteBar := createFavoriteGrid(myWindow, *user)
	spacer := canvas.NewText("", color.White)
	title := canvas.NewText("Groupie Trackers", color.White)
	title.TextSize = 42
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}
	subtitle := canvas.NewText(" Découvrez de nouveaux artistes : ", color.White)
	subtitle.TextSize = 16
	subtitle2 := canvas.NewText(" Découvrez les dernières sorties : ", color.White)
	subtitle2.TextSize = 16
	subtitle3 := canvas.NewText(" Redécouvrez le meilleur des vieux albums : ", color.White)
	subtitle3.TextSize = 16
	subtitle4 := canvas.NewText(" Vos favoris : ", color.White)
	subtitle4.TextSize = 16

	favoriteButton := widget.NewButton("Mes favoris", func() {
		FavoritePage(myApp)
		myWindow.Hide()
	})

	favoriteButton.Importance = widget.HighImportance

	favorite := container.NewHBox(subtitle4, layout.NewSpacer(), favoriteButton, spacer)
	backgroundRect := canvas.NewRectangle(color.RGBA{R: 16, G: 16, B: 16, A: 255}) //

	content := container.NewVBox(spacer, spacer, title, spacer, spacer,
		subtitle, spacer, rdmBar, spacer, favorite, spacer, favoriteBar, spacer, subtitle2, spacer, lastAlbumBar, spacer, subtitle3, spacer, firstAlbumBar,
		spacer)

	scrollContainer := container.NewVScroll(content)

	myWindow.SetContent(container.NewBorder(navBar, nil, nil, nil, backgroundRect, scrollContainer))

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}

func FavoritePage(myApp fyne.App) {
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
	myWindow.SetContent(container.NewBorder(navBar, nil, nil, nil, scrollContainer))
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}

var PasswordChange = false

func AccountPage(myApp fyne.App) {
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

	if user.Ppf == "" {
		form := container.NewVBox(
			spacer,
			spacer,
			subtitle,
			spacer,
			changePasswordButton,
			changePpfButton,
		)

		content := container.NewVBox(
			title,
			form,
		)

		centeredContent := container.NewCenter(content)
		myWindow.SetContent(container.NewBorder(navBar, nil, nil, nil, centeredContent))
	} else {

		ppf := loadImageFromURL(user.Ppf) // JARRIVE PAS A REDIMENTIONNER LIMAGE
		form := container.NewVBox(
			spacer,
			changePasswordButton,
			changePpfButton,
		)

		content := container.NewVBox(
			title,
			ppf,
			subtitle,
			form,
		)

		centeredContent := container.NewCenter(content)
		myWindow.SetContent(container.NewBorder(navBar, nil, nil, nil, centeredContent))
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

var PpfChange = false

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
