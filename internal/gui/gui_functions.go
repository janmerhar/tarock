package gui

import (
	"fyne.io/fyne/v2"
)

func Load_image(img_path string) fyne.Resource {
	icon, _ := fyne.LoadResourceFromPath(img_path)
	return icon
}

func Create_game_window(a fyne.App) fyne.Window {
	w := a.NewWindow("Solitaire")
	w.SetPadded(false)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(1280, 720))

	return w
}
