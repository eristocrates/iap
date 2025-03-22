package windows

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func MainWindow(a fyne.App) {
	title := "Main Screen"
	w := a.NewWindow(title)

	w.Resize(fyne.NewSize(900, 475))
	w.SetContent(widget.NewButton("Hello World Screen", func() {
		HelloWindow(a)
	}))

	w.Show()
}
