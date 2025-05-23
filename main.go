package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"apple2dos-gui/ui"
)

func main() {
	a := app.New()
	w := a.NewWindow("Apple ][ DOS 3.3")

	terminal := ui.NewTerminal()
	w.SetContent(container.NewVBox(
		terminal.View,
	))

	w.Resize(600, 400)
	w.ShowAndRun()
}
