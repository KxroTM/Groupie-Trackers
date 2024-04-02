package app

import (
	"Groupie_Trackers/go/functions"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ConcertMap(artist functions.Artist, myApp fyne.App) {
	myWindow := MyApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)

	navBar := createNavBar(myWindow)

	backButton := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		ConcertPage(artist, myApp)
		myWindow.Hide()
	})

	topContent := container.NewVBox(
		container.NewHBox(
			backButton,
			layout.NewSpacer(),
			widget.NewLabel("Localisation du Concert :"),
			layout.NewSpacer(),
		),
	)

	content := container.NewStack(container.NewBorder(container.NewVBox(navBar, topContent), nil, nil, nil))

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})
	myWindow.SetContent(content)
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}
