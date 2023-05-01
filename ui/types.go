package ui

import (
	"fyne.io/fyne/v2"
	"github.com/jmarron7/go-pixl/apptype"
	"github.com/jmarron7/go-pixl/pxcanvas"
	"github.com/jmarron7/go-pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
