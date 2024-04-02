package app

import (
	"Groupie_Trackers/go/functions"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ConcertPage(artist functions.Artist, myApp fyne.App) {
	myWindow := MyApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)

	navBar := createNavBar(myWindow)

	locations := functions.LocationsData()
	dates := functions.DatesData()

	var artistlocations []string
	var artistdates []string

	for _, location := range locations.Index {
		if artist.ID == location.ID {
			artistlocations = location.Locations
		}
	}

	for _, date := range dates.Index {
		if artist.ID == date.ID {
			artistdates = date.Dates
		}
	}

	fmt.Println(artistlocations, artistdates)
	fmt.Println(len(artistlocations), len(artistdates))

	// Création de la planification des concerts pour l'artiste
	concertSchedule := GetConcertSchedule(artistlocations, artistdates)

	// Affichage de la planification des concerts
	fmt.Println(concertSchedule)

	// Création du conteneur VBox pour afficher les informations
	creationDateRange := container.NewVBox(
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewLabel("Concerts :"),
			layout.NewSpacer(),
		),
	)

	// Itération sur la planification des concerts pour créer les éléments d'interface graphique
	for location, dates := range concertSchedule {
		hbox := container.NewHBox(
			layout.NewSpacer(),
			widget.NewLabel(location), // Ajout du label de l'emplacement
		)

		// Ajout des dates à la HBox
		for _, date := range dates {
			hbox.Add(widget.NewLabel(date))
		}

		// Espace flexible à droite pour aligner les dates au centre
		hbox.Add(layout.NewSpacer())

		// Ajout de la HBox au conteneur VBox
		creationDateRange.Add(hbox)
	}

	content := container.NewStack(container.NewBorder(navBar, nil, nil, nil, creationDateRange))

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})
	myWindow.SetContent(content)
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}

// Fonction pour obtenir toutes les dates pour chaque emplacement
func GetConcertSchedule(locations, dates []string) map[string][]string {
	concertSchedule := make(map[string][]string)
	currentLocation := 0

	for i := 0; i < len(dates); i++ {
		if strings.HasPrefix(dates[i], "*") {
			if i != 0 {
				if strings.HasPrefix(dates[i-1], "*") {
					currentLocation++
				}
			}
			concertSchedule[locations[currentLocation]] = append(concertSchedule[locations[currentLocation]], dates[i])

		} else {
			concertSchedule[locations[currentLocation]] = append(concertSchedule[locations[currentLocation]], dates[i])
			if i != len(dates)-1 {
				if strings.HasPrefix(dates[i+1], "*") {
					currentLocation++
				}
			}
		}
	}
	return concertSchedule
}
