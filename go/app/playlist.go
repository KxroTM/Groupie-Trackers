package app

import (
	"Groupie_Trackers/go/functions"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

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

	createOneButton := widget.NewButton("Créer une playlist", func() {
		CreatePlaylistPage(myApp, functions.Artist{})
		myWindow.Hide()
	})
	buttonContent := container.NewHBox(layout.NewSpacer(), createOneButton)

	playlistsContainer := container.NewVBox()

	for _, playlist := range user.Playlists.Playlist {
		playlistTemp := playlist

		Title := canvas.NewText(" "+playlist.Name, color.White)
		Title.TextSize = 22
		Title.TextStyle = fyne.TextStyle{Bold: true}
		playlistButton := widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
			OnePlaylistPage(myApp, user, playlistTemp.Name)
			myWindow.Hide()
		})
		deleteButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
			functions.DeletePlaylist(user.Username, playlistTemp.Name)
			PlaylistPage(myApp)
			myWindow.Hide()
		})

		titleContent := container.NewHBox(Title, layout.NewSpacer(), deleteButton, playlistButton, spacer)
		playlistGrid = createPlaylistGrid(myWindow, *user, playlist.Name)
		playlistsContainer.Add(container.NewVBox(titleContent, spacer, playlistGrid, spacer))
	}

	fullContent := container.NewVScroll(container.NewVBox(spacer, buttonContent, title, spacer, playlistsContainer, spacer))
	content := container.NewStack(container.NewBorder(navBar, nil, nil, nil, BackgroundRect, fullContent))

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})
	myWindow.SetContent(content)
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}

func OnePlaylistPage(myApp fyne.App, user *functions.Account, playlist string) {
	myWindow := myApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)
	user = functions.UserBuild(user.Username)
	navBar := createNavBar(myWindow)

	backButton := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		PlaylistPage(myApp)
		myWindow.Hide()
	})

	playlistGrid := createFullPlaylistGrid(myWindow, user, playlist)
	gridContainer := container.NewStack()
	gridContainer.Add(playlistGrid)

	content := container.NewStack(container.NewBorder(container.NewVBox(navBar, container.NewHBox(backButton, layout.NewSpacer())),
		nil, nil, nil, BackgroundRect, gridContainer))

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
		if artist.Name != "" {
			playlistCreate = true
			ArtistPage(artist, myApp)
			myWindow.Hide()
		} else {
			PlaylistPage(myApp)
			myWindow.Hide()
		}
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

	myWindow.SetContent(container.NewBorder(container.NewVBox(navBar), nil, nil, nil, BackgroundRect, centeredContent))
	myWindow.Resize(fyne.NewSize(600, 400))
	myWindow.CenterOnScreen()

	myWindow.Show()
}
