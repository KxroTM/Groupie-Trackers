package app

import (
	"Groupie_Trackers/go/functions"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// struct to store an image and its URL
type ImageEntry struct {
	URL   string
	Image *canvas.Image
}

var imageCache []*ImageEntry

// load an image from a URL, caching it for future use
func loadImageFromURL(url string) *canvas.Image {
	for _, imageEntry := range imageCache {
		if imageEntry.URL == url {
			return imageEntry.Image
		}
	}

	image := downloadImage(url)
	imageEntry := &ImageEntry{URL: url, Image: image}
	imageCache = append(imageCache, imageEntry)

	return image
}

// download an image from a URL
func downloadImage(url string) *canvas.Image {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to download image:", err)
		return canvas.NewImageFromImage(image.NewRGBA(image.Rect(0, 0, 1, 1)))
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read image data:", err)
		return canvas.NewImageFromImage(image.NewRGBA(image.Rect(0, 0, 1, 1)))
	}

	tmpFile, err := os.CreateTemp("", "image-*.png")
	if err != nil {
		fmt.Println("Failed to create a temp file for the image:", err)
		return canvas.NewImageFromImage(image.NewRGBA(image.Rect(0, 0, 1, 1)))
	}
	defer tmpFile.Close()

	_, err = tmpFile.Write(data)
	if err != nil {
		fmt.Println("Failed to write image data to temp file:", err)
		return canvas.NewImageFromImage(image.NewRGBA(image.Rect(0, 0, 1, 1)))
	}

	image := canvas.NewImageFromFile(tmpFile.Name())
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(200, 200))

	return image
}

func createArtistsGrid(artists []functions.Artist, w fyne.Window) fyne.CanvasObject {
	var artistCards []fyne.CanvasObject

	for _, artist := range artists {
		artistTemp := artist
		image := loadImageFromURL(artist.Image)
		image.FillMode = canvas.ImageFillContain
		button := widget.NewButton(artist.Name, func() {
			ArtistPage(artistTemp, MyApp)
			w.Hide()
		})
		card := container.NewVBox(image, button)
		artistCards = append(artistCards, card)
	}

	grid := container.NewGridWithColumns(4, artistCards...)
	scrollContainer := container.NewVScroll(grid)

	return scrollContainer
}

func createNavBar(myWindow fyne.Window) *fyne.Container {
	homeButton := widget.NewButton("Accueil", func() {
		HomePage(MyApp)
		myWindow.Hide()
	})

	accountButton := widget.NewButton("Mon compte", func() {
		AccountPage(MyApp)
		myWindow.Hide()
	})

	favoriteButton := widget.NewButton("Favoris", func() {
		FavoritePage(MyApp)
		myWindow.Hide()
	})

	logoutButton := widget.NewButtonWithIcon("", theme.LogoutIcon(), func() {
		functions.LogOut()
		LoginPage(MyApp)
		myWindow.Hide()
	})

	researchButton := widget.NewButtonWithIcon("", theme.SearchIcon(), func() {
		SearchPage(MyApp)
		myWindow.Hide()
	})

	space := canvas.NewText("  ", color.Transparent)
	space2 := canvas.NewText("      ", color.Transparent)

	content := container.NewHBox(space, researchButton, layout.NewSpacer(), homeButton,
		accountButton, favoriteButton, layout.NewSpacer(), logoutButton, space)

	return container.NewVBox(space2, content, space2)
}

func createSearchBar(artists []functions.Artist, gridContainer *fyne.Container, w fyne.Window) fyne.CanvasObject {
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher un artiste...")

	searchEntry.OnChanged = func(text string) {
		filteredArtists := functions.Search(artists, text)
		newGrid := createArtistsGrid(filteredArtists, w)
		gridContainer.Objects = []fyne.CanvasObject{newGrid}
		gridContainer.Refresh()
	}

	return searchEntry
}

// Create a grid with 4 random artists
func createRandomArtistsGrid(w fyne.Window) fyne.CanvasObject {
	var artistCards []fyne.CanvasObject
	var artists []functions.Artist
	artistsContent := functions.ArtistData()
	check1 := false
	check2 := false
	check3 := false

	for i := 0; i < len(artistsContent); i++ {
		if len(artists) == 4 {
			break
		}
		artistsTemp := artistsContent[rand.Intn(len(artistsContent))]
		artistsTemp2 := artistsContent[rand.Intn(len(artistsContent))]
		artistsTemp3 := artistsContent[rand.Intn(len(artistsContent))]

		for i := 0; i < len(artists); i++ {
			if artistsTemp.Name == artists[i].Name {
				check1 = true
			}
			if artistsTemp2.Name == artists[i].Name {
				check2 = true
			}
			if artistsTemp3.Name == artists[i].Name {
				check3 = true
			}
		}
		if !check1 {
			artists = append(artists, artistsTemp)
		} else if !check2 {
			artists = append(artists, artistsTemp2)
		} else if !check3 {
			artists = append(artists, artistsTemp3)
		}

		if i == len(artistsContent)-1 && len(artists) < 4 {
			artists = append(artists, artistsContent[0])
		}
	}

	for _, artist := range artists {
		artistTemp := artist
		image, _ := fyne.LoadResourceFromURLString(artist.Image)

		img := canvas.NewImageFromResource(image)

		img.SetMinSize(fyne.NewSize(200, 200))
		btn := widget.NewButton(" ", func() {
			ArtistPage(artistTemp, MyApp)
			w.Hide()
		})

		container1 := container.New(
			layout.NewStackLayout(),
			btn,
			widget.NewCard("", "  "+artist.Name, img),
		)

		card := container.NewVBox(container1)
		artistCards = append(artistCards, card)
	}

	grid := container.NewGridWithColumns(4, artistCards...)

	return grid
}

// Create a grid with 4 artists by a list of artists
func createCustomArtistsGrid(w fyne.Window, artistContent functions.AllArtists) fyne.CanvasObject {
	var artistCards []fyne.CanvasObject
	var artists []functions.Artist

	for i := 0; i < 4; i++ {
		artists = append(artists, artistContent[i])
	}

	for _, artist := range artists {
		artistTemp := artist
		image, _ := fyne.LoadResourceFromURLString(artist.Image)

		img := canvas.NewImageFromResource(image)

		img.SetMinSize(fyne.NewSize(200, 200))
		btn := widget.NewButton(" ", func() {
			ArtistPage(artistTemp, MyApp)
			w.Hide()
		})

		container1 := container.New(
			layout.NewStackLayout(),
			btn,
			widget.NewCard("", "  "+artist.Name, img),
		)

		card := container.NewVBox(container1)
		artistCards = append(artistCards, card)
	}

	grid := container.NewGridWithColumns(4, artistCards...)

	return grid
}

// Create a grid with the 4 first favorites artists
func createFavoriteGrid(w fyne.Window, user functions.Account) fyne.CanvasObject {
	var artistCards []fyne.CanvasObject
	var artists []functions.Artist

	if len(user.Favorites) == 0 {
		spacer := canvas.NewText("", color.White)
		card := container.NewVBox(spacer, spacer, spacer, spacer, spacer, spacer, spacer, spacer)
		return card

	}

	artistContent := functions.ArtistData()

	for i := 0; i < len(user.Favorites) && i < 4; i++ {
		for j := 0; j < len(artistContent); j++ {
			if user.Favorites[i] == artistContent[j].Name {
				artists = append(artists, artistContent[j])
			}
		}
	}

	for _, artist := range artists {
		artistTemp := artist
		image, _ := fyne.LoadResourceFromURLString(artist.Image)

		img := canvas.NewImageFromResource(image)

		img.SetMinSize(fyne.NewSize(200, 200))
		btn := widget.NewButton(" ", func() {
			ArtistPage(artistTemp, MyApp)
			w.Hide()
		})

		container1 := container.New(
			layout.NewStackLayout(),
			btn,
			widget.NewCard("", "  "+artist.Name, img),
		)

		card := container.NewVBox(container1)
		artistCards = append(artistCards, card)
	}

	grid := container.NewGridWithColumns(4, artistCards...)

	return grid
}

// Create a grid with all the favorite artists
func createAllFavoriteGrid(w fyne.Window, user functions.Account) fyne.CanvasObject {
	var artistCards []fyne.CanvasObject
	var artists []functions.Artist

	if len(user.Favorites) == 0 {
		text := canvas.NewText("Votre liste de favori est vide.", color.White)
		return text
	}

	artistContent := functions.ArtistData()

	for i := 0; i < len(user.Favorites); i++ {
		for j := 0; j < len(artistContent); j++ {
			if user.Favorites[i] == artistContent[j].Name {
				artists = append(artists, artistContent[j])
			}
		}
	}

	for _, artist := range artists {
		artistTemp := artist
		image, _ := fyne.LoadResourceFromURLString(artist.Image)

		img := canvas.NewImageFromResource(image)

		img.SetMinSize(fyne.NewSize(200, 200))
		btn := widget.NewButton(" ", func() {
			ArtistPage(artistTemp, MyApp)
			w.Hide()
		})

		container1 := container.New(
			layout.NewStackLayout(),
			btn,
			widget.NewCard("", "  "+artist.Name, img),
		)

		card := container.NewVBox(container1)
		artistCards = append(artistCards, card)
	}

	grid := container.NewGridWithColumns(4, artistCards...)

	return grid
}

// Create a grid with the 4 first history artists
func createHistoryGrid(w fyne.Window, user functions.Account) fyne.CanvasObject {
	var artistCards []fyne.CanvasObject
	var artists []functions.Artist

	if len(user.History) == 0 {
		spacer := canvas.NewText("", color.White)
		card := container.NewVBox(spacer, spacer, spacer, spacer, spacer, spacer, spacer, spacer)
		return card
	}

	artistContent := functions.ArtistData()

	// Get the last 4 artists in the history
	for i := len(user.History) - 1; i >= 0 && len(artists) < 4; i-- {
		for j := 0; j < len(artistContent); j++ {
			if len(artists) == 0 {
				if user.History[i] == artistContent[j].Name {
					artists = append(artists, artistContent[j])
				}
			} else {
				if user.History[i] != artists[len(artists)-1].Name {
					if user.History[i] == artistContent[j].Name {
						artists = append(artists, artistContent[j])
					}
				}
			}
		}
	}

	for _, artist := range artists {
		artistTemp := artist
		image, _ := fyne.LoadResourceFromURLString(artist.Image)

		img := canvas.NewImageFromResource(image)

		img.SetMinSize(fyne.NewSize(200, 200))
		btn := widget.NewButton(" ", func() {
			ArtistPage(artistTemp, MyApp)
			w.Hide()
		})

		container1 := container.New(
			layout.NewStackLayout(),
			btn,
			widget.NewCard("", "  "+artist.Name, img),
		)

		card := container.NewVBox(container1)
		artistCards = append(artistCards, card)
	}

	grid := container.NewGridWithColumns(4, artistCards...)

	return grid
}

// Create a playlist grid with the 4 first playlists
func createPlaylistGrid(w fyne.Window, user functions.Account, playlist string) fyne.CanvasObject {
	var artistCards []fyne.CanvasObject
	var artists []functions.Artist
	var index int

	for i := 0; i < len(user.Playlists.Playlist); i++ {
		if user.Playlists.Playlist[i].Name == playlist {
			index = i
			break
		}
	}

	if len(user.Playlists.Playlist[index].Songs) == 0 {
		spacer := canvas.NewText("", color.White)
		card := container.NewVBox(spacer, spacer, spacer, spacer, spacer, spacer, spacer, spacer)
		return card
	}

	artistContent := functions.ArtistData()

	for i := 0; i < len(user.Playlists.Playlist[index].Songs) && i < 4; i++ {
		for j := 0; j < len(artistContent); j++ {
			if user.Playlists.Playlist[index].Songs[i] == artistContent[j].Name {
				artists = append(artists, artistContent[j])
			}
		}
	}

	for _, artist := range artists {
		artistTemp := artist
		image, _ := fyne.LoadResourceFromURLString(artist.Image)

		img := canvas.NewImageFromResource(image)

		img.SetMinSize(fyne.NewSize(200, 200))
		btn := widget.NewButton(" ", func() {
			ArtistPage(artistTemp, MyApp)
			w.Hide()
		})

		container1 := container.New(
			layout.NewStackLayout(),
			btn,
			widget.NewCard("", "  "+artist.Name, img),
		)

		card := container.NewVBox(container1)
		artistCards = append(artistCards, card)
	}

	grid := container.NewGridWithColumns(4, artistCards...)

	return grid
}

// Create a playlist grid with all the artists in the playlist
func createFullPlaylistGrid(w fyne.Window, user *functions.Account, playlist string) fyne.CanvasObject {
	var artistCards []fyne.CanvasObject
	var artists []functions.Artist
	var index int

	for i := 0; i < len(user.Playlists.Playlist)-1; i++ {
		if user.Playlists.Playlist[i].Name == playlist {
			index = i
			break
		}
	}

	if len(user.Playlists.Playlist[index].Songs) == 0 {
		spacer := canvas.NewText("", color.White)
		card := container.NewVBox(spacer, spacer, spacer, spacer, spacer, spacer, spacer, spacer)
		return card
	}

	artistContent := functions.ArtistData()

	for i := 0; i < len(user.Playlists.Playlist[index].Songs); i++ {
		for j := 0; j < len(artistContent); j++ {
			if user.Playlists.Playlist[index].Songs[i] == artistContent[j].Name {
				artists = append(artists, artistContent[j])
			}
		}
	}

	for _, artist := range artists {
		artistTemp := artist
		image := loadImageFromURL(artist.Image)
		image.FillMode = canvas.ImageFillContain
		image.SetMinSize(fyne.NewSize(200, 200))
		button := widget.NewButton(artist.Name, func() {
			ArtistPage(artistTemp, MyApp)
			w.Hide()
		})
		deleteButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
			functions.DeleteSongFromPlaylist(user.Username, playlist, artistTemp.Name)
			OnePlaylistPage(MyApp, user, playlist)
			w.Hide()
		})
		buttonContainer := container.NewHBox(button, deleteButton)
		card := container.NewVBox(image, buttonContainer)
		artistCards = append(artistCards, card)
	}

	grid := container.NewGridWithColumns(4, artistCards...)
	scrollContainer := container.NewVScroll(grid)

	return scrollContainer
}

// Load an image from a file and encode
func loadImageBinary(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Create a custom icon by a local image file
func createCustomIcon(filename string) fyne.Resource {
	image, _ := loadImageBinary(filename)
	myImage := fyne.NewStaticResource("my_image.png", image)
	return myImage
}

// Create a fake spotify embed
func SpotifyEmbed(artist functions.Artist) *fyne.Container {
	var date *canvas.Text
	var date2 *canvas.Text
	var embed *fyne.Container
	var sliderValue float64

	img := loadImageFromURL(artist.Image)
	img.SetMinSize(fyne.NewSize(150, 150))
	title := canvas.NewText(artist.Name, color.White)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.TextSize = 20

	if len(artist.Members) > 4 {
		members := ""
		members2 := ""

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

		date = canvas.NewText(artist.FirstAlbum+", "+members, color.White)
		date2 = canvas.NewText(members2, color.White)
	} else {
		artistMembers := ""
		for i := 0; i < len(artist.Members); i++ {
			artistMembers += artist.Members[i]
			if i != len(artist.Members)-1 {
				artistMembers += ", "
			}
		}
		date = canvas.NewText(artist.FirstAlbum+", "+artistMembers, color.White)
	}

	spacer := canvas.NewText("  ", color.White)
	borderSpacer := canvas.NewText("", color.Transparent)
	spacertxt := canvas.NewText("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", color.Transparent)

	musicTime := 190.0

	s := widget.NewSlider(0, musicTime)

	var isPlaying = true
	var playButton *widget.Button

	timeLeftTxt := canvas.NewText("03:10", color.White)
	currentTimetxt := canvas.NewText("00:00", color.White)

	playButton = widget.NewButtonWithIcon("", theme.MediaPauseIcon(), func() {
		if isPlaying {
			playButton.SetIcon(theme.MediaPlayIcon())
			isPlaying = false
		} else {
			playButton.SetIcon(theme.MediaPauseIcon())
			isPlaying = true

			go func() {
				for {
					if !isPlaying {
						break
					}
					for i := sliderValue; i <= musicTime; i++ {
						if !isPlaying {
							break
						}
						time.Sleep(time.Second)
						s.Value = i
						s.Refresh()

						currentSeconds := int(s.Value) % 60
						currentMinutes := int(s.Value) / 60
						currentTime := fmt.Sprintf("%02d:%02d", currentMinutes, currentSeconds)
						currentTimetxt.Text = currentTime
						currentTimetxt.Refresh()

						remainingSeconds := int(musicTime-s.Value) % 60
						remainingMinutes := int(musicTime-s.Value) / 60
						remainingTime := fmt.Sprintf("%02d:%02d", remainingMinutes, remainingSeconds)
						timeLeftTxt.Text = remainingTime
						timeLeftTxt.Refresh()
					}
				}
			}()
		}
	})

	go func() {
		for {
			if !isPlaying {
				break
			}
			for i := 0.0; i <= musicTime; i++ {
				if !isPlaying {
					break
				}
				time.Sleep(time.Second)
				s.Value = i
				s.Refresh()
				sliderValue = s.Value

				currentSeconds := int(s.Value) % 60
				currentMinutes := int(s.Value) / 60
				currentTime := fmt.Sprintf("%02d:%02d", currentMinutes, currentSeconds)
				currentTimetxt.Text = currentTime
				currentTimetxt.Refresh()

				remainingSeconds := int(musicTime-s.Value) % 60
				remainingMinutes := int(musicTime-s.Value) / 60
				remainingTime := fmt.Sprintf("%02d:%02d", remainingMinutes, remainingSeconds)
				timeLeftTxt.Text = remainingTime
				timeLeftTxt.Refresh()
			}
		}
	}()

	rightContent := container.NewHBox(timeLeftTxt, spacer, spacer, spacer, playButton)
	slider := container.NewBorder(nil, nil, currentTimetxt, rightContent, s)

	if len(artist.Members) > 4 {
		embed = container.NewVBox(spacer, title, date, date2, spacertxt, slider, spacertxt)
	} else {
		embed = container.NewVBox(spacer, title, date, spacertxt, slider, spacertxt)
	}

	body := container.NewHBox(borderSpacer, img, spacer, embed)
	bg := canvas.NewRectangle(color.RGBA{4, 59, 92, 255})
	img.Resize(fyne.NewSize(180, 180))

	return container.NewBorder(nil, nil, nil, nil, bg, body)
}
