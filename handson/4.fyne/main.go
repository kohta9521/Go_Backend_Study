package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("hello")
	w.SetContent(
		widget.NewLabel("Hello Fyne!"),
	)

	w.ShowAndRun()
}