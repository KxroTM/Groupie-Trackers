package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func ConcertMap(myApp fyne.App) {
	myWindow := MyApp.NewWindow("Groupie Trackers")
	myWindow.SetIcon(Icon)

	navBar := createNavBar(myWindow)
	content := container.NewStack(container.NewBorder(navBar, nil, nil, nil))

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})
	myWindow.SetContent(content)
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.Show()
}
