package main

import (
	"fyne.io/fyne/v2/app"

	"iap/internal/adapters/uidrivers/fyne/windows"
)

func main() {
	a := app.New()
	windows.MainWindow(a)
	a.Run()
}
