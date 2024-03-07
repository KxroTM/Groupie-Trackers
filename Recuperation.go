package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"net/http"
)

type Artist struct {
	Name     string `json:"name"`
	ImageURL string `json:"image"`
}

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
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to download image:", err)
		return canvas.NewImageFromFile("placeholder.png")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read image data:", err)
		return canvas.NewImageFromFile("placeholder.png")
	}

	tmpFile, err := ioutil.TempFile("", "image-*.png")
	if err != nil {
		fmt.Println("Failed to create a temp file for the image:", err)
		return canvas.NewImageFromFile("placeholder.png")
	}
	defer tmpFile.Close()

	_, err = tmpFile.Write(data)
	if err != nil {
		fmt.Println("Failed to write image data to temp file:", err)
		return canvas.NewImageFromFile("placeholder.png")
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

	return container.NewHBox(layout.NewSpacer(), homeButton, aboutButton, contactButton, layout.NewSpacer())
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hip Hop Showcase")

	navBar := createNavBar()

	artists, err := fetchArtists()
	if err != nil {
		fmt.Println("Erreur lors de la récupération des artistes:", err)
		return
	}

	artistsGrid := createArtistsGrid(artists)
	mainContent := container.NewBorder(navBar, nil, nil, nil, artistsGrid)

	myWindow.SetContent(mainContent)
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
