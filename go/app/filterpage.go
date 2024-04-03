package app

import (
	"Groupie_Trackers/go/functions"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func FilterPage(myApp fyne.App) {
	myWindow := myApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)
	gridContainer := container.NewGridWithColumns(1)

	navBar := createNavBar(myWindow)
	researchButton := widget.NewButton("Recherche sans Filtre", func() {
		SearchPage(MyApp)
		myWindow.Hide()
	})

	// Fonction pour mettre à jour le Label avec l'année correspondant à la valeur du slider
	updateLabelYear := func(label *widget.Label, value float64) {
		year := int(value)
		label.SetText(strconv.Itoa(year))
	}

	// Créer des Labels pour afficher les années
	labelCreationDateStart := widget.NewLabel("1958")
	labelCreationDateEnd := widget.NewLabel("2015")

	sliderCreationDateStart := widget.NewSlider(1958, 2015) // 1958 for Bee Gees and 2015 for Juice Wrld
	sliderCreationDateEnd := widget.NewSlider(1958, 2015)   // 1958 for Bee Gees and 2015 for Juice Wrld

	sliderCreationDateStart.SetValue(1958)
	sliderCreationDateEnd.SetValue(2015)

	// Mettre à jour les Labels à chaque fois que la valeur des sliders change
	sliderCreationDateStart.OnChanged = func(value float64) {
		if value > sliderCreationDateEnd.Value {
			// Si la valeur de début dépasse la valeur de fin, ajustez la valeur de fin
			sliderCreationDateEnd.SetValue(value)
		}
		updateLabelYear(labelCreationDateStart, value)
	}
	sliderCreationDateEnd.OnChanged = func(value float64) {
		if value < sliderCreationDateStart.Value {
			// Si la valeur de fin est inférieure à la valeur de début, ajustez la valeur de début
			sliderCreationDateStart.SetValue(value)
		}
		updateLabelYear(labelCreationDateEnd, value)
	}

	creationDateRange := container.NewVBox(
		widget.NewLabel("Date de Création :"),
		container.NewHBox(
			layout.NewSpacer(),
			sliderCreationDateStart,
			labelCreationDateStart,
			layout.NewSpacer(),
		),
		container.NewHBox(
			layout.NewSpacer(),
			sliderCreationDateEnd,
			labelCreationDateEnd,
			layout.NewSpacer(),
		),
	)

	// Déclaration et initialisation des sliders pour First Album Date Range

	labelFirstAlbumStart := widget.NewLabel("1967")
	labelFirstAlbumDateEnd := widget.NewLabel("2018")

	// Créer les sliders
	sliderFirstAlbumStart := widget.NewSlider(1967, 2018)
	sliderFirstAlbumEnd := widget.NewSlider(1967, 2018)

	sliderFirstAlbumStart.SetValue(1967)
	sliderFirstAlbumEnd.SetValue(2018)

	// Mettre à jour les Labels à chaque fois que la valeur des sliders change
	sliderFirstAlbumStart.OnChanged = func(value float64) {
		if value > sliderFirstAlbumEnd.Value {
			// Si la valeur de début dépasse la valeur de fin, ajustez la valeur de fin
			sliderFirstAlbumEnd.SetValue(value)
		}
		updateLabelYear(labelFirstAlbumStart, value)
	}
	sliderFirstAlbumEnd.OnChanged = func(value float64) {
		if value < sliderFirstAlbumStart.Value {
			// Si la valeur de fin est inférieure à la valeur de début, ajustez la valeur de début
			sliderFirstAlbumStart.SetValue(value)
		}
		updateLabelYear(labelFirstAlbumDateEnd, value)
	}

	firstAlbumRange := container.NewVBox(
		widget.NewLabel("Publication du Premier Album :"),
		container.NewHBox(
			layout.NewSpacer(),
			sliderFirstAlbumStart,
			labelFirstAlbumStart,
			layout.NewSpacer(),
		),
		container.NewHBox(
			layout.NewSpacer(),
			sliderFirstAlbumEnd,
			labelFirstAlbumDateEnd,
			layout.NewSpacer(),
		),
	)

	// Déclaration et initialisation de l'entry pour Number of Members
	oneM := widget.NewCheck("1", func(checked bool) {})
	twoM := widget.NewCheck("2", func(checked bool) {})
	threeM := widget.NewCheck("3", func(checked bool) {})
	fourM := widget.NewCheck("4", func(checked bool) {})
	fiveM := widget.NewCheck("5", func(checked bool) {})
	sixM := widget.NewCheck("6", func(checked bool) {})
	sevenM := widget.NewCheck("7", func(checked bool) {})

	numMembers := container.NewVBox(
		widget.NewLabel("Nombre de Membres :"),
		container.NewVBox(
			container.NewHBox(
				layout.NewSpacer(),
				oneM,
				twoM,
				threeM,
				fourM,
				layout.NewSpacer(),
			),
			container.NewHBox(
				layout.NewSpacer(),
				fiveM,
				sixM,
				sevenM,
				layout.NewSpacer(),
			),
		),
	)
	oneM.SetChecked(true)
	twoM.SetChecked(true)
	threeM.SetChecked(true)
	fourM.SetChecked(true)
	fiveM.SetChecked(true)
	sixM.SetChecked(true)
	sevenM.SetChecked(true)

	DateRange := container.NewHBox(
		layout.NewSpacer(),
		creationDateRange,
		layout.NewSpacer(),
		firstAlbumRange,
		layout.NewSpacer(),
		numMembers,
		layout.NewSpacer(),
	)

	// Déclaration et initialisation des checkboxes pour Locations
	countriesname := []string{"germany", "saudi_arabia", "netherlands_antilles", "argentina", "australia",
		"austria", "belgium", "belarus", "brazil", "canada", "chile", "china", "colombia",
		"south_korea", "costa_rica", "denmark", "united_arab_emirates", "usa", "spain",
		"finland", "france", "greece", "hungary", "india", "indonesia", "ireland", "italy",
		"japan", "mexico", "norway", "new_caledonia", "new_zealand", "netherlands", "peru",
		"philippines", "poland", "french_polynesia", "portugal", "qatar", "romania", "uk",
		"slovakia", "sweden", "switzerland", "taiwan", "czechia", "thailand",
	}

	germany := widget.NewCheck("Allemagne", func(checked bool) {})
	saudi_arabia := widget.NewCheck("Arabie saoudite", func(checked bool) {})
	netherlands_antilles := widget.NewCheck("Antilles Néerlandaises", func(checked bool) {})
	argentina := widget.NewCheck("Argentine", func(checked bool) {})
	australia := widget.NewCheck("Australie", func(checked bool) {})
	austria := widget.NewCheck("Autriche", func(checked bool) {})
	belgium := widget.NewCheck("Belgique", func(checked bool) {})
	belarus := widget.NewCheck("Biélorussie", func(checked bool) {})
	brazil := widget.NewCheck("Brésil", func(checked bool) {})
	canada := widget.NewCheck("Canada", func(checked bool) {})
	chile := widget.NewCheck("Chili", func(checked bool) {})
	china := widget.NewCheck("Chine", func(checked bool) {})
	colombia := widget.NewCheck("Colombie", func(checked bool) {})
	south_korea := widget.NewCheck("Corée du Sud", func(checked bool) {})
	costa_rica := widget.NewCheck("Costa Rica", func(checked bool) {})
	denmark := widget.NewCheck("Danemark", func(checked bool) {})
	united_arab_emirates := widget.NewCheck("Émirats Arabes Unis", func(checked bool) {})
	usa := widget.NewCheck("États-Unis", func(checked bool) {})
	spain := widget.NewCheck("Espagne", func(checked bool) {})
	finland := widget.NewCheck("Finlande", func(checked bool) {})
	france := widget.NewCheck("France", func(checked bool) {})
	greece := widget.NewCheck("Grèce", func(checked bool) {})
	hungary := widget.NewCheck("Hongrie", func(checked bool) {})
	india := widget.NewCheck("Inde", func(checked bool) {})
	indonesia := widget.NewCheck("Indonésie", func(checked bool) {})
	ireland := widget.NewCheck("Irlande", func(checked bool) {})
	italy := widget.NewCheck("Italie", func(checked bool) {})
	japan := widget.NewCheck("Japon", func(checked bool) {})
	mexico := widget.NewCheck("Mexique", func(checked bool) {})
	norway := widget.NewCheck("Norvège", func(checked bool) {})
	new_caledonia := widget.NewCheck("Nouvelle-Calédonie", func(checked bool) {})
	new_zealand := widget.NewCheck("Nouvelle-Zélande", func(checked bool) {})
	netherlands := widget.NewCheck("Pays-Bas", func(checked bool) {})
	peru := widget.NewCheck("Pérou", func(checked bool) {})
	philippines := widget.NewCheck("Philippines", func(checked bool) {})
	poland := widget.NewCheck("Pologne", func(checked bool) {})
	french_polynesia := widget.NewCheck("Polynésie Française", func(checked bool) {})
	portugal := widget.NewCheck("Portugal", func(checked bool) {})
	qatar := widget.NewCheck("Qatar", func(checked bool) {})
	romania := widget.NewCheck("Roumanie", func(checked bool) {})
	uk := widget.NewCheck("Royaume-Uni", func(checked bool) {})
	slovakia := widget.NewCheck("Slovaquie", func(checked bool) {})
	sweden := widget.NewCheck("Suède", func(checked bool) {})
	switzerland := widget.NewCheck("Suisse", func(checked bool) {})
	taiwan := widget.NewCheck("Taïwan", func(checked bool) {})
	czechia := widget.NewCheck("Tchéquie", func(checked bool) {})
	thailand := widget.NewCheck("Thaïlande", func(checked bool) {})

	countrieswidget := []*widget.Check{
		germany, saudi_arabia, netherlands_antilles, argentina, australia,
		austria, belgium, belarus, brazil, canada, chile, china, colombia,
		south_korea, costa_rica, denmark, united_arab_emirates, usa, spain,
		finland, france, greece, hungary, india, indonesia, ireland, italy,
		japan, mexico, norway, new_caledonia, new_zealand, netherlands, peru,
		philippines, poland, french_polynesia, portugal, qatar, romania, uk,
		slovakia, sweden, switzerland, taiwan, czechia, thailand,
	}

	countriesUncheck := widget.NewButton("Uncheck All Countries", func() {
		uncheckChecks(
			germany, saudi_arabia, netherlands_antilles, argentina, australia,
			austria, belgium, belarus, brazil, canada, chile, china, colombia,
			south_korea, costa_rica, denmark, united_arab_emirates, usa, spain,
			finland, france, greece, hungary, india, indonesia, ireland, italy,
			japan, mexico, norway, new_caledonia, new_zealand, netherlands, peru,
			philippines, poland, french_polynesia, portugal, qatar, romania, uk,
			slovakia, sweden, switzerland, taiwan, czechia, thailand,
		)
	})

	countriesCheck := widget.NewButton("Check All Countries", func() {
		checkChecks(
			germany, saudi_arabia, netherlands_antilles, argentina, australia,
			austria, belgium, belarus, brazil, canada, chile, china, colombia,
			south_korea, costa_rica, denmark, united_arab_emirates, usa, spain,
			finland, france, greece, hungary, india, indonesia, ireland, italy,
			japan, mexico, norway, new_caledonia, new_zealand, netherlands, peru,
			philippines, poland, french_polynesia, portugal, qatar, romania, uk,
			slovakia, sweden, switzerland, taiwan, czechia, thailand,
		)
	})

	locations := container.NewVBox(
		container.NewVBox(
			container.NewHBox(
				layout.NewSpacer(),
				widget.NewLabel("Localisation des concerts :"),
				layout.NewSpacer(),
			),
			container.NewHBox(
				layout.NewSpacer(),
				countriesUncheck,
				countriesCheck,
				layout.NewSpacer(),
			),
		),
		container.NewHBox(
			layout.NewSpacer(),
			germany,
			saudi_arabia,
			netherlands_antilles,
			argentina,
			australia,
			austria,
			belgium,
			layout.NewSpacer(),
		),
		container.NewHBox(
			layout.NewSpacer(),
			belarus,
			brazil,
			canada,
			chile,
			china,
			colombia,
			south_korea,
			costa_rica,
			layout.NewSpacer(),
		),
		container.NewHBox(
			layout.NewSpacer(),
			denmark,
			united_arab_emirates,
			usa,
			spain,
			finland,
			france,
			greece,
			hungary,
			layout.NewSpacer(),
		),
		container.NewHBox(
			layout.NewSpacer(),
			india,
			indonesia,
			ireland,
			italy,
			japan,
			mexico,
			norway,
			new_caledonia,
			layout.NewSpacer(),
		),
		container.NewHBox(
			layout.NewSpacer(),
			new_zealand,
			netherlands,
			peru,
			philippines,
			poland,
			french_polynesia,
			portugal,
			layout.NewSpacer(),
		),
		container.NewHBox(
			layout.NewSpacer(),
			qatar,
			romania,
			uk,
			slovakia,
			sweden,
			switzerland,
			taiwan,
			czechia,
			thailand,
			layout.NewSpacer(),
		),
	)

	applyButton := widget.NewButton("Apply Filters", func() {
		artists := functions.ArtistData()

		artists = functions.ArtistbyCreationDateRange(artists, sliderCreationDateStart.Value, sliderCreationDateEnd.Value)

		// Convert float64 values to strings
		artists = functions.ArtistbyFirstAlbumDateRange(artists, sliderFirstAlbumStart.Value, sliderFirstAlbumEnd.Value)

		// Obtenir les nombres de membres cochés
		checkedNumbers := getCheckedNumbers(oneM, twoM, threeM, fourM, fiveM, sixM, sevenM)
		// Appliquer le filtre sur le nombre de membres
		artists = functions.ArtistbyNumberofMemberCheck(artists, checkedNumbers)

		// Obtenir les pays cochés
		checkedcountries := getCheckedCountries(countrieswidget, countriesname)
		// Appliquer le filtre sur les pays cochés
		LocationsData := functions.LocationsData()
		artists = functions.ArtistbyCountry(artists, LocationsData, checkedcountries)

		artistsGrid := createArtistsGrid(artists, myWindow)

		// Remplacez le contenu du gridContainer par le nouveau artistsGrid
		gridContainer.Objects = []fyne.CanvasObject{artistsGrid}
		gridContainer.Refresh()
	})

	resetButton := widget.NewButton("Reset Filters", func() {
		// Ici, vous pouvez maintenant accéder directement aux variables
		sliderCreationDateStart.SetValue(1958)
		sliderCreationDateEnd.SetValue(2015)
		sliderFirstAlbumStart.SetValue(1967)
		sliderFirstAlbumEnd.SetValue(2018)

		oneM.SetChecked(true)
		twoM.SetChecked(true)
		threeM.SetChecked(true)
		fourM.SetChecked(true)
		fiveM.SetChecked(true)
		sixM.SetChecked(true)
		sevenM.SetChecked(true)

	})

	filterContainer := container.NewVBox(
		DateRange,
		locations,
		container.NewHBox(
			layout.NewSpacer(),
			applyButton,
			resetButton,
			layout.NewSpacer(),
		),
	)

	botContent := container.NewVScroll(gridContainer) // Placer gridContainer dans un conteneur défilable

	topContent := container.NewVBox(navBar, researchButton, filterContainer)

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})

	myWindow.SetContent(container.NewBorder(topContent, nil, nil, nil, botContent)) // Utiliser scrollContainer à la place de gridContainer
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 850))
	myWindow.SetFixedSize(true)
	myWindow.Show()
}

func uncheckChecks(checks ...*widget.Check) {
	for _, check := range checks {
		check.SetChecked(false)
	}
}

func checkChecks(checks ...*widget.Check) {
	for _, check := range checks {
		check.SetChecked(true)
	}
}

func getCheckedNumbers(checks ...*widget.Check) []int {
	var checkedNumbers []int
	for i, check := range checks {
		if check.Checked {
			checkedNumbers = append(checkedNumbers, i+1) // Ajouter 1 car les nombres de membres commencent à partir de 1
		}
	}
	return checkedNumbers
}

func getCheckedCountries(widgets []*widget.Check, countriesname []string) []string {
	var checkedCountries []string

	for i, widget := range widgets {
		// i += 1
		if widget.Checked {
			checkedCountries = append(checkedCountries, countriesname[i])
		}
	}

	return checkedCountries
}
