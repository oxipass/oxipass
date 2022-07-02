package main

import (
	"fyne.io/fyne/v2"
	"github.com/oxipass/oxilib"
	"log"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (oxiApp *OxiApp) addItemForm() fyne.CanvasObject {
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),

		widget.NewToolbarAction(theme.ContentClearIcon(), func() {
			oxiApp.Window.SetContent(oxiApp.mainScreen())
		}),
	)

	iconNames := oxilib.GetFontAwesomeList()
	iconNames = iconNames[0:20]
	itemNameEntry := widget.NewEntry()
	itemNameEntry.SetPlaceHolder("item name")
	itemIconEntry := widget.NewSelectEntry(iconNames)
	itemIconEntry.SetPlaceHolder("icon name")
	backButton := widget.NewButtonWithIcon("Back", theme.NavigateBackIcon(), func() {
		oxiApp.Window.SetContent(oxiApp.mainScreen())
	})
	addButton := widget.NewButtonWithIcon("Add", theme.ConfirmIcon(), func() {
		log.Println("Add item")
		var item oxilib.UpdateItemForm
		item.Icon = itemIconEntry.Text
		item.Name = itemNameEntry.Text
		_, err := oxilib.GetInstance().AddNewItem(item)
		if err != nil {
			dialog.ShowError(err, oxiApp.Window)
		}
		oxiApp.Window.SetContent(oxiApp.mainScreen())
	})

	vBox := container.NewVBox(
		itemNameEntry,
		itemIconEntry,
		addButton,
		backButton,
	)

	content := container.New(layout.NewBorderLayout(toolbar, nil, nil, nil), toolbar, vBox)
	return content
}
