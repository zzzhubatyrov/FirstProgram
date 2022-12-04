package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("MyFirstProg")
	w.Resize(fyne.NewSize(550, 550))

	// Компонент CARD

	card1 := widget.NewCard(
		"petrfact",
		"Снимаю видео про сервера и программирование",
		canvas.NewImageFromFile("./img/rpm.jpg"),
	)

	card2 := widget.NewCard(
		"whqhub",
		"Info",
		widget.NewButton("Subscribe", func() {}),
	)

	box := container.NewHBox(card1, card2)

	w.SetContent(box)
	w.ShowAndRun()
}
