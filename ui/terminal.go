package ui

import (
	"apple2dos-gui/dos"
	"fyne.io/fyne/v2/widget"
)

type Terminal struct {
	View *widget.Entry
}

func NewTerminal() *Terminal {
	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("APPLE ][ DOS 3.3\nREADY\nDOS> ")

	output.OnSubmitted = func(command string) {
		result := dos.Execute(command)
		output.SetText(output.Text + "\n" + result + "\nDOS> ")
	}

	return &Terminal{View: output}
}
