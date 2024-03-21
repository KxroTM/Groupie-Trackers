package app

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"io"
	"net/http"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Artist struct {
	Name     string `json:"name"`
	ImageURL string `json:"image"`
}

type ImageEntry struct {
	URL   string
	Image *canvas.Image
}

var imageCache []*ImageEntry

func fetchArtists() ([]Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching artists: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return nil, fmt.Errorf("error decoding API response: %v", err)
	}

	return artists, nil
}

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

func createArtistsGrid(artists []Artist) fyne.CanvasObject {
	var artistCards []fyne.CanvasObject

	for _, artist := range artists {
		image := loadImageFromURL(artist.ImageURL)
		label := widget.NewLabel(artist.Name)
		label.Alignment = fyne.TextAlignCenter // Assurez l'alignement du texte au centre sous l'image

		card := container.NewVBox(image, label)
		artistCards = append(artistCards, card)
	}

	grid := container.NewGridWithColumns(4, artistCards...)
	scrollContainer := container.NewVScroll(grid)

	return scrollContainer
}

func createNavBar() *fyne.Container {
	homeButton := widget.NewButton("Accueil", nil)
	aboutButton := widget.NewButton("À Propos", nil)
	contactButton := widget.NewButton("Contact", nil)
	text := canvas.NewText("Welcome "+user.Username, color.White)
	space := canvas.NewText(text.Text, color.Transparent)
	space2 := canvas.NewText("      ", color.Transparent)

	return container.NewHBox(layout.NewSpacer(), space, space2, homeButton, aboutButton, contactButton, layout.NewSpacer(), text, space2)
}

func createSearchBar(artists []Artist, gridContainer *fyne.Container) fyne.CanvasObject {
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher un artiste...")

	searchEntry.OnChanged = func(text string) {
		filteredArtists := filterArtists(artists, text)
		newGrid := createArtistsGrid(filteredArtists)
		gridContainer.Objects = []fyne.CanvasObject{newGrid}
		gridContainer.Refresh()
	}

	return searchEntry
}

// Fonction pour filtrer les artistes basé sur la saisie de l'utilisateur
func filterArtists(artists []Artist, filterText string) []Artist {
	var filtered []Artist
	for _, artist := range artists {
		if filterText == "" || strings.HasPrefix(strings.ToLower(artist.Name), strings.ToLower(filterText)) {
			filtered = append(filtered, artist)
		}
	}
	return filtered
}
