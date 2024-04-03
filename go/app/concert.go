package app

import (
	"Groupie_Trackers/go/functions"
	"fmt"
	"os/exec"
	"runtime"
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
			datebox.Add(
				widget.NewButton(CodeToShowDates(date), func() {
					slicedate, err := functions.DateStringToStringSlice(date)
					if err != nil {
						fmt.Println("Error:", err)
					}
					openURL("https://calendar.google.com/calendar/u/0/r/day/" + slicedate[2] + "/" + slicedate[1] + "/" + slicedate[0])
				}))
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
	city = translateFrenchCity(city)

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

func translateFrenchCity(city string) string {
	cities := map[string]string{
		"NORTH CAROLINA":         "CAROLINE DU NORD",
		"GEORGIA":                "GÉORGIE",
		"LOS ANGELES":            "LOS ANGELES",
		"SAITAMA":                "SAITAMA",
		"OSAKA":                  "OSAKA",
		"NAGOYA":                 "NAGOYA",
		"PENROSE":                "PENROSE",
		"DUNEDIN":                "DUNEDIN",
		"PLAYA DEL CARMEN":       "PLAYA DEL CARMEN",
		"PAPEETE":                "PAPEETE",
		"NOUMEA":                 "NOUMÉA",
		"LONDON":                 "LONDRES",
		"LAUSANNE":               "LAUSANNE",
		"LYON":                   "LYON",
		"VICTORIA":               "VICTORIA",
		"NEW SOUTH WALES":        "NOUVELLE-GALLES DU SUD",
		"QUEENSLAND":             "QUEENSLAND",
		"AUCKLAND":               "AUCKLAND",
		"YOGYAKARTA":             "YOGYAKARTA",
		"BRATISLAVA":             "BRATISLAVA",
		"BUDAPEST":               "BUDAPEST",
		"MINSK":                  "MINSK",
		"CALIFORNIA":             "CALIFORNIE",
		"NEVADA":                 "NEVADA",
		"SAO PAULO":              "SAO PAULO",
		"SAN ISIDRO":             "SAN ISIDRO",
		"ARIZONA":                "ARIZONA",
		"TEXAS":                  "TEXAS",
		"STOCKHOLM":              "STOCKHOLM",
		"WERCHTER":               "WERCHTER",
		"LISBON":                 "LISBONNE",
		"BILBAO":                 "BILBAO",
		"BOGOTA":                 "BOGOTA",
		"NEW YORK":               "NEW YORK",
		"DUSSELDORF":             "DÜSSELDORF",
		"AARHUS":                 "AARHUS",
		"MANCHESTER":             "MANCHESTER",
		"FRANKFURT":              "FRANCFORT",
		"BERLIN":                 "BERLIN",
		"COPENHAGEN":             "COPENHAGUE",
		"DOHA":                   "DOHA",
		"MINNESOTA":              "MINNESOTA",
		"ILLINOIS":               "ILLINOIS",
		"MUMBAI":                 "MUMBAI",
		"ABU DHABI":              "ABOU DHABI",
		"PENNSYLVANIA":           "PENNSYLVANIE",
		"WESTCLIFF ON SEA":       "WESTCLIFF-ON-SEA",
		"MERKERS":                "MERKERS",
		"MAINE":                  "MAINE",
		"GOTHENBURG":             "GOTHENBURG",
		"FLORIDA":                "FLORIDE",
		"SOUTH CAROLINA":         "CAROLINE DU SUD",
		"PAGNEY DERRIERE BARINE": "PAGNEY-DERRIÈRE-BARINE",
		"HAMBURG":                "HAMBURG",
		"BOULOGNE BILLANCOURT":   "BOULOGNE-BILLANCOURT",
		"NIMES":                  "NÎMES",
		"SION":                   "SION",
		"OSTRAVA":                "OSTRAVA",
		"KLAGENFURT":             "KLAGENFURT",
		"FREYMING MERLEBACH":     "FREYMING-MERLEBACH",
		"ZARAGOZA":               "ZARAGOZA",
		"MADRID":                 "MADRID",
		"BARCELONA":              "BARCELONE",
		"RIO DE JANEIRO":         "RIO DE JANEIRO",
		"RECIFE":                 "RECIFE",
		"LEIPZIG":                "LEIPZIG",
		"SALEM":                  "SALEM",
		"MONCHENGLADBACH":        "MONCHENGLADBACH",
		"CUXHAVEN":               "CUXHAVEN",
		"SKANDERBORG":            "SKANDERBORG",
		"AMSTERDAM":              "AMSTERDAM",
		"BURRIANA":               "BURRIANA",
		"OULU":                   "OULU",
		"NAPOCA":                 "NAPOCA",
		"RIYADH":                 "RIYAD",
		"CANTON":                 "CANTON",
		"QUEBEC":                 "QUÉBEC",
		"LAS VEGAS":              "LAS VEGAS",
		"MEXICO CITY":            "MEXICO",
		"MONTERREY":              "MONTERREY",
		"DEL MAR":                "DEL MAR",
		"WASHINGTON":             "WASHINGTON",
		"WEST MELBOURNE":         "WEST MELBOURNE",
		"PARIS":                  "PARIS",
		"MISSOURI":               "MISSOURI",
		"CHICAGO":                "CHICAGO",
		"BIRMINGHAM":             "BIRMINGHAM",
		"SYDNEY":                 "SYDNEY",
		"MADISON":                "MADISON",
		"TORONTO":                "TORONTO",
		"CLEVELAND":              "CLEVELAND",
		"BOSTON":                 "BOSTON",
		"UTAH":                   "UTAH",
		"GLASGOW":                "GLASGOW",
		"DUBLIN":                 "DUBLIN",
		"CARDIFF":                "CARDIFF",
		"ABERDEEN":               "ABERDEEN",
		"WARSAW":                 "VARSOVIE",
		"MILAN":                  "MILAN",
		"MICHIGAN":               "MICHIGAN",
		"NEW HAMPSHIRE":          "NEW HAMPSHIRE",
		"SOCHAUX":                "SOCHAUX",
		"EINDHOVEN":              "EINDHOVEN",
		"OSLO":                   "OSLO",
		"COLORADO":               "COLORADO",
		"JAKARTA":                "JAKARTA",
		"HUIZHOU":                "HUIZHOU",
		"CHANGZHOU":              "CHANGZHOU",
		"HONG KONG":              "HONG KONG",
		"SANYA":                  "SANYA",
		"AALBORG":                "AALBORG",
		"SEATTLE":                "SEATTLE",
		"OMAHA":                  "OMAHA",
		"KANSAS CITY":            "KANSAS CITY",
		"ST LOUIS":               "ST LOUIS",
		"INDIANAPOLIS":           "INDIANAPOLIS",
		"ROSEMONT":               "ROSEMONT",
		"GRAND RAPIDS":           "GRAND RAPIDS",
		"MONTREAL":               "MONTRÉAL",
		"NEWARK":                 "NEWARK",
		"UNIONDALE":              "UNIONDALE",
		"PHILADELPHIA":           "PHILADELPHIE",
		"HERSHEY":                "HERSHEY",
		"PITTSBURGH":             "PITTSBURGH",
		"COLUMBIA":               "COLUMBIA",
		"SANTIAGO":               "SANTIAGO",
		"HOUSTON":                "HOUSTON",
		"ATLANTA":                "ATLANTA",
		"NEW ORLEANS":            "NOUVELLE-ORLÉANS",
		"FRAUENFELD":             "FRAUENFELD",
		"TURKU":                  "TURKU",
		"BROOKLYN":               "BROOKLYN",
		"IMOLA":                  "IMOLA",
		"VIENNA":                 "VIENNE",
		"KRAKOW":                 "CRACOVIE",
		"ZURICH":                 "ZURICH",
		"AMITYVILLE":             "AMITYVILLE",
		"MINNEAPOLIS":            "MINNEAPOLIS",
		"DETROIT":                "DÉTROIT",
		"OAKLAND":                "OAKLAND",
		"CHARLOTTE":              "CHARLOTTE",
		"INGLEWOOD":              "INGLEWOOD",
		"WINDSOR":                "WINDSOR",
		"CINCINNATI":             "CINCINNATI",
		"ANAHEIM":                "ANAHEIM",
		"MANILA":                 "MANILLE",
		"BRISBANE":               "BRISBANE",
		"MELBOURNE":              "MELBOURNE",
		"LIMA":                   "LIMA",
		"GRONINGEN":              "GRONINGEN",
		"ANTWERP":                "ANVERS",
		"PICO RIVERA":            "PICO RIVERA",
		"BERWYN":                 "BERWYN",
		"DALLAS":                 "DALLAS",
		"BRIXTON":                "BRIXTON",
		"ROTSELAAR":              "ROTSELAAR",
		"ALABAMA":                "ALABAMA",
		"MASSACHUSETTS":          "MASSACHUSETTS",
		"ATHENS":                 "ATHÈNES",
		"FLORENCE":               "FLORENCE",
		"LANDGRAAF":              "LANDGRAAF",
		"BURSWOOD":               "BURSWOOD",
		"WELLINGTON":             "WELLINGTON",
		"SEVILLE":                "SÉVILLE",
		"BANGKOK":                "BANGKOK",
		"TAIPEI":                 "TAÏPEI",
		"SEOUL":                  "SÉOUL",
		"MUNICH":                 "MUNICH",
		"MANNHEIM":               "MANNHEIM",
		"SAN FRANCISCO":          "SAN FRANCISCO",
		"BUENOS AIRES":           "BUENOS AIRES",
		"PORTO ALEGRE":           "PORTO ALEGRE",
		"BELO HORIZONTE":         "BELO HORIZONTE",
		"LA PLATA":               "LA PLATA",
		"DUBAI":                  "DUBAÏ",
		"WILLEMSTAD":             "WILLEMSTAD",
		"BRASILIA":               "BRASILIA",
		"OKLAHOMA":               "OKLAHOMA",
		"SCHEESSEL":              "SCHEESSEL",
		"ST GALLEN":              "SAINT-GALL",
		"GDYNIA":                 "GDYNIA",
		"ARRAS":                  "ARRAS",
		"SAN JOSE":               "SAN JOSÉ",
		"NICKELSDORF":            "NICKELSDORF",
		"OREGON":                 "OREGON",
		"VANCOUVER":              "VANCOUVER",
		"PRAGUE":                 "PRAGUE",
	}

	// Convertit le nom du pays entré en majuscules pour correspondre aux clés de la map
	if translatedCountry, ok := cities[city]; ok {
		return translatedCountry
	}
	// Renvoie le pays non trouvé tel quel ou un message spécifique
	return city
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
		"UK":                   "ROYAUME-UNI",
		"SLOVAKIA":             "SLOVAQUIE",
		"SWEDEN":               "SUÈDE",
		"SWITZERLAND":          "SUISSE",
		"TAIWAN":               "TAÏWAN",
		"CZECHIA":              "TCHÉQUIE",
		"THAILAND":             "THAÏLANDE",
	}

	// Convertit le nom du pays entré en majuscules pour correspondre aux clés de la map
	if translatedCountry, ok := invertedCountryTranslations[country]; ok {
		return translatedCountry
	}
	// Renvoie le pays non trouvé tel quel ou un message spécifique
	return country
}

func openURL(url string) error {
	var cmd *exec.Cmd

	// Détermine le système d'exploitation en cours
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}

	// Exécute la commande pour ouvrir le navigateur
	err := cmd.Start()
	if err != nil {
		return err
	}
	return nil
}

func containsString(slice []string, element string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}
