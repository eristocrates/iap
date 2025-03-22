package windows

// https://docs.fyne.io/started/hello
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func HelloWindow(a fyne.App) {
	w := a.NewWindow("Hello World")
	w.SetContent(widget.NewLabel("Hello World!"))
	w.Show()
}
