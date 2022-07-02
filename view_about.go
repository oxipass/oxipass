package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/bykovme/gotrans"
)

func (oxiApp *OxiApp) buildAboutView() fyne.CanvasObject {
	aboutView := container.NewVBox(
		widget.NewLabel(""),
		container.NewCenter(oxiApp.Logo),
		container.NewCenter(widget.NewLabel(gotrans.T("version")+" "+cAppVersion)),
		layout.NewSpacer(),
	)
	return aboutView
}
