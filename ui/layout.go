package ui

import "fyne.io/fyne/v2/container"

func Setup(app *AppInit) {
	SetupMenus(app)
	swatchesContainer := BuildSwatches(app)
	coloPicker := SetupColorPicker(app)

	appLayout := container.NewBorder(nil, swatchesContainer, nil, coloPicker, app.PixlCanvas)
	app.PixlWindow.SetContent(appLayout)
}
