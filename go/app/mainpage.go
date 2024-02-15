package app

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Mainpage() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie-Trackers")

	title := widget.NewLabel("Mainpage")
	content := container.NewVBox(
		title,
	)

	myWindow.SetContent(content)
	myWindow.SetMaster()

	myWindow.ShowAndRun()
}
