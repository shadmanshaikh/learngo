package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new application
	myApp := app.New()

	// Create a new window
	myWindow := myApp.NewWindow("Hello, Fyne!")

	// Set the content of the window
	label := widget.NewLabel("Welcome to Fyne!")
	button := widget.NewButton("Click Me", func() {
		label.SetText("Button Clicked!")
	})

	// Arrange the widgets in a vertical box
	content := container.NewVBox(label, button)
	myWindow.SetContent(content)

	// Show and run the application
	myWindow.ShowAndRun()
}
