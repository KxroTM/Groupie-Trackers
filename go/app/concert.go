package app

import (
	"Groupie_Trackers/go/functions"
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
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

	backButton := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		ArtistPage(artist, myApp)
		myWindow.Hide()
	})

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

	// Création de la planification des concerts pour l'artiste
	concertSchedule := GetConcertSchedule(artistlocations, artistdates)

	// Création du conteneur VBox pour afficher les informations
	topContent := container.NewVBox(
		container.NewHBox(
			backButton,
			layout.NewSpacer(),
			widget.NewLabel("Concerts :"),
			layout.NewSpacer(),
		),
	)

	concerts := container.NewVBox()

	// Itération sur la planification des concerts pour créer les éléments d'interface graphique
	for location, dates := range concertSchedule {
		printable_location := CodeToShowLocation(location)

		locationButton := widget.NewButton(printable_location, func() {
			ConcertMap(artist, myApp, printable_location)
			myWindow.Hide()
		})
		hbox := container.NewHBox(
			locationButton,
		)
		hbox.Add(layout.NewSpacer())

		datebox := container.NewHBox()
		for _, date := range dates {
			datebox.Add(widget.NewLabel(CodeToShowDates(date)))
		}
		scrollContainer := container.NewHScroll(datebox)

		allbox := container.NewVBox(
			hbox,
			scrollContainer,
		)

		concerts.Add(allbox)
	}

	content := container.NewStack(container.NewBorder(container.NewVBox(navBar, topContent), nil, nil, nil, concerts))
	scrollContent := container.NewVScroll(content)

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})
	myWindow.SetContent(scrollContent)
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

func CodeToShowLocation(location string) string {
	Splitlocation := strings.Split(location, "-")

	city := Splitlocation[0]
	city = strings.ReplaceAll(city, "_", " ")
	city = strings.ToUpper(city)

	country := Splitlocation[1]
	country = strings.ReplaceAll(country, "_", " ")
	country = strings.ToUpper(country)
	country = translateFrenchCountries(country)

	return city + ", " + country
}

func CodeToShowDates(date string) string {
	if date[0] == '*' {
		date = date[1:]
	}
	slicedate, err := functions.DateStringToIntSlice(date)
	if err != nil {
		fmt.Println("Error:", err)
	}

	printabledate := time.Date(slicedate[2], time.Month(slicedate[1]), slicedate[0], 0, 0, 0, 0, time.UTC)

	formattedDate := printabledate.Format("Mon. 2 Jan. 2006")
	formattedDate = translateFrenchFormattedDate(formattedDate)

	return formattedDate

}

func translateFrenchFormattedDate(date string) string {
	// Remplacer les abréviations des noms de jours et de mois en anglais par les abréviations en français
	frenchDays := map[string]string{
		"Mon.": "Lun.",
		"Tue.": "Mar.",
		"Wed.": "Mer.",
		"Thu.": "Jeu.",
		"Fri.": "Ven.",
		"Sat.": "Sam.",
		"Sun.": "Dim.",
	}

	frenchMonths := map[string]string{
		"Jan.": "Janv.",
		"Feb.": "Févr.",
		"Mar.": "Mars",
		"Apr.": "Avr.",
		"May":  "Mai",
		"Jun.": "Juin",
		"Jul.": "Juil.",
		"Aug.": "Août",
		"Sep.": "Sept.",
		"Oct.": "Oct.",
		"Nov.": "Nov.",
		"Dec.": "Déc.",
	}

	for english, french := range frenchDays {
		date = strings.ReplaceAll(date, english, french)
	}

	for english, french := range frenchMonths {
		date = strings.ReplaceAll(date, english, french)
	}

	return date
}

func translateFrenchCountries(country string) string {
	invertedCountryTranslations := map[string]string{
		"GERMANY":              "ALLEMAGNE",
		"SAUDI ARABIA":         "ARABIE SAOUDITE",
		"NETHERLANDS ANTILLES": "ANTILLES NÉERLANDAISES",
		"ARGENTINA":            "ARGENTINE",
		"AUSTRALIA":            "AUSTRALIE",
		"AUSTRIA":              "AUTRICHE",
		"BELGIUM":              "BELGIQUE",
		"BELARUS":              "BIÉLORUSSIE",
		"BRAZIL":               "BRÉSIL",
		"CANADA":               "CANADA",
		"CHILE":                "CHILI",
		"CHINA":                "CHINE",
		"COLOMBIA":             "COLOMBIE",
		"SOUTH KOREA":          "CORÉE DU SUD",
		"COSTA RICA":           "COSTA RICA",
		"DENMARK":              "DANEMARK",
		"UNITED ARAB EMIRATES": "ÉMIRATS ARABES UNIS",
		"USA":                  "ÉTATS-UNIS",
		"SPAIN":                "ESPAGNE",
		"FINLAND":              "FINLANDE",
		"FRANCE":               "FRANCE",
		"GREECE":               "GRÈCE",
		"HUNGARY":              "HONGRIE",
		"INDIA":                "INDE",
		"INDONESIA":            "INDONÉSIE",
		"IRELAND":              "IRLANDE",
		"ITALY":                "ITALIE",
		"JAPAN":                "JAPON",
		"MEXICO":               "MEXIQUE",
		"NORWAY":               "NORVÈGE",
		"NEW CALEDONIA":        "NOUVELLE-CALÉDONIE",
		"NEW ZEALAND":          "NOUVELLE-ZÉLANDE",
		"NETHERLANDS":          "PAYS-BAS",
		"PERU":                 "PÉROU",
		"PHILIPPINES":          "PHILIPPINES",
		"POLAND":               "POLOGNE",
		"FRENCH POLYNESIA":     "POLYNÉSIE FRANÇAISE",
		"PORTUGAL":             "PORTUGAL",
		"QATAR":                "QATAR",
		"ROMANIA":              "ROUMANIE",
		"UNITED KINGDOM":       "ROYAUME-UNI",
		"SLOVAKIA":             "SLOVAQUIE",
		"SWEDEN":               "SUÈDE",
		"SWITZERLAND":          "SUISSE",
		"TAIWAN":               "TAÏWAN",
		"CZECH REPUBLIC":       "TCHÉQUIE",
		"THAILAND":             "THAÏLANDE",
	}

	// Convertit le nom du pays entré en majuscules pour correspondre aux clés de la map
	if translatedCountry, ok := invertedCountryTranslations[country]; ok {
		return translatedCountry
	}
	// Renvoie le pays non trouvé tel quel ou un message spécifique
	return "Pays non trouvé"
}
