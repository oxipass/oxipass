package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type OxiApp struct {
	Window        fyne.Window
	App           fyne.App
	CurrentSearch string

	ItemsList          *widget.List
	ItemsListContainer *ItemsContainer

	FieldsList          *widget.List
	FieldsListContainer *FieldsContainer

	TagsList          *widget.List
	TagsListContainer *TagsContainer

	ItemNameLabel     *widget.Label
	MainScreenContent *fyne.CanvasObject

	SplitContainer *container.Split

	Logo *canvas.Image
}

func (oxiApp *OxiApp) Init() {
	//canvas.NewImageFromReader()
	oxiApp.Logo = canvas.NewImageFromFile("icons/Icon.png")
	size := fyne.NewSize(100, 100)
	oxiApp.Logo.SetMinSize(size)
	oxiApp.Logo.Resize(size)

}

func (oxiApp *OxiApp) OnClose() {
	oxiApp.App.Preferences().SetFloat(cPrefSplitOffsetKey, oxiApp.SplitContainer.Offset)
	size := oxiApp.Window.Content().Size()
	oxiApp.App.Preferences().SetInt(cPrefWidth, int(size.Width))
	oxiApp.App.Preferences().SetInt(cPrefHeight, int(size.Height))
}
