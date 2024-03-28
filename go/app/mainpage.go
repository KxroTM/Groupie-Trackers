package app

import (
	"Groupie_Trackers/go/functions"
	"fmt"
	"image"
	"image/color"
	"io"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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

func createArtistsGrid(artists []functions.Artist) fyne.CanvasObject {
	var artistCards []fyne.CanvasObject

	for _, artist := range artists {
		image := loadImageFromURL(artist.Image)
		label := widget.NewLabel(artist.Name)
		label.Alignment = fyne.TextAlignCenter // Assurez l'alignement du texte au centre sous l'image

		card := container.NewVBox(image, label)
		artistCards = append(artistCards, card)
	}

	grid := container.NewGridWithColumns(4, artistCards...)
	scrollContainer := container.NewVScroll(grid)

	return scrollContainer
}

func createNavBar(myWindow fyne.Window) *fyne.Container {
	homeButton := widget.NewButton("Accueil", func() {
		LoginPage(MyApp)
		myWindow.Hide()
	})
	aboutButton := widget.NewButton("À Propos", func() {
		Propospage()
		myWindow.Hide()
	})
	contactButton := widget.NewButton("Contact", func() {
		Contactpage()
		myWindow.Hide()
	})

	text := canvas.NewText("Welcome "+user.Username, color.Black)
	space := canvas.NewText(text.Text, color.Transparent)
	space2 := canvas.NewText("      ", color.Transparent)

	return container.NewHBox(layout.NewSpacer(), space, space2, homeButton, aboutButton, contactButton, layout.NewSpacer(), text, space2)
}

func createSearchBar(artists []functions.Artist, gridContainer *fyne.Container) fyne.CanvasObject {
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher un artiste...")

	searchEntry.OnChanged = func(text string) {
		filteredArtists := functions.Search(artists, text)
		newGrid := createArtistsGrid(filteredArtists)
		gridContainer.Objects = []fyne.CanvasObject{newGrid}
		gridContainer.Refresh()
	}

	return searchEntry
}
