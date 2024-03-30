package app

import (
	"Groupie_Trackers/go/functions"
	"fmt"
	"image"
	"image/color"
	"io"
	"math/rand"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ImageEntry struct {
	URL   string
	Image *canvas.Image
}

var imageCache []*ImageEntry

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

	return grid
}

func createCustomArtistsGrid(w fyne.Window, artistContent functions.AllArtists) fyne.CanvasObject {
	var artistCards []fyne.CanvasObject
	var artists []functions.Artist

	for i := 0; i < 4; i++ {
		artists = append(artists, artistContent[i])
	}

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

	return grid
}

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

	return grid
}

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

	return grid
}
