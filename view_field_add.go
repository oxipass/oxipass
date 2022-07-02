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

func (oxiApp *OxiApp) addFieldForm(itemId int64) fyne.CanvasObject {
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),

		widget.NewToolbarAction(theme.ContentClearIcon(), func() {
			oxiApp.Window.SetContent(oxiApp.mainScreen())
		}),
	)

	iconNames := oxilib.GetFontAwesomeList()
	iconNames = iconNames[0:20]
	itemNameEntry := widget.NewEntry()
	itemIconEntry := widget.NewSelectEntry(iconNames)
	valueTypes := oxilib.GetValueTypes()
	valueTypeEntry := widget.NewSelectEntry(valueTypes)
	valueEntry := widget.NewEntry()
	backButton := widget.NewButtonWithIcon("Back", theme.NavigateBackIcon(), func() {
		oxiApp.Window.SetContent(oxiApp.mainScreen())
	})
	addButton := widget.NewButtonWithIcon("Add", theme.ConfirmIcon(), func() {
		log.Println("Add field")
		var field oxilib.UpdateFieldForm
		field.Icon = itemIconEntry.Text
		field.Name = itemNameEntry.Text
		field.ItemID = itemId
		field.ValueType = valueTypeEntry.Text
		field.Value = valueEntry.Text
		_, err := oxilib.GetInstance().AddNewField(field)
		if err != nil {
			dialog.ShowError(err, oxiApp.Window)
		}
		oxiApp.Window.SetContent(oxiApp.mainScreen())
	})

	vBox := container.NewVBox(
		itemNameEntry,
		itemIconEntry,
		valueTypeEntry,
		valueEntry,
		addButton,
		backButton,
	)

	content := container.New(layout.NewBorderLayout(toolbar, nil, nil, nil), toolbar, vBox)
	return content
}
