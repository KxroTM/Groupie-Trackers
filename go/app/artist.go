package app

import (
	"Groupie_Trackers/go/functions"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

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
