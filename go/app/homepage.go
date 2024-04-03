package app

import (
	"Groupie_Trackers/go/functions"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	myWindow.Resize(fyne.NewSize(900, 650))
	myWindow.SetFixedSize(true)
	myWindow.Show()
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
	myWindow.Resize(fyne.NewSize(900, 650))
	myWindow.SetFixedSize(true)
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
	myWindow.Resize(fyne.NewSize(900, 650))
	myWindow.SetFixedSize(true)
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
