package main

import (
	"fyne.io/fyne/v2"
	"github.com/oxipass/oxilib"
	"log"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/bykovme/gotrans"
)

func (oxiApp *OxiApp) searchOnChanged(searchTerm string) {
	oxiApp.CurrentSearch = searchTerm
	log.Println("Search term: " + searchTerm)
}

func (oxiApp *OxiApp) mainScreen() fyne.CanvasObject {
	//link, err := url.Parse("https://bykov.me")
	//if err != nil {
	//	fyne.LogError("Could not parse URL", err)
	//}
	oxiInstance := oxilib.GetInstance()

	if oxiInstance.IsNew() == true {
		return oxiApp.createRegisterForm()
	} else if oxiInstance.IsLocked() == true {
		return oxiApp.createLoginForm()
	}

	toolbar := oxiApp.createMainToolbar()

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder(gotrans.T("search"))
	searchEntry.SetText(oxiApp.CurrentSearch)
	oxiApp.ItemNameLabel = widget.NewLabelWithStyle("Select item to show fields", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	searchEntry.OnChanged = oxiApp.searchOnChanged

	if oxiApp.FieldsListContainer == nil {
		oxiApp.FieldsListContainer = CreateFieldsList()
	}

	if oxiApp.ItemsListContainer == nil {
		oxiApp.ItemsListContainer, _ = CreateItemsList()
	} else {
		_ = UpdateItemsList(oxiApp.ItemsListContainer)
	}

	if oxiApp.FieldsList == nil {
		oxiApp.FieldsList = widget.NewList(oxiApp.FieldsListContainer.GetListSize,
			oxiApp.FieldsListContainer.GetListDesign,
			oxiApp.FieldsListContainer.UpdateItem)
		oxiApp.FieldsList.OnSelected = oxiApp.FieldsListContainer.OnSelected
	} else {
		if oxiApp.ItemsListContainer != nil {
			_ = UpdateFieldsList(oxiApp.FieldsListContainer, oxiApp.ItemsListContainer.selectedItem)
			oxiApp.ItemNameLabel.SetText(oxiApp.ItemsListContainer.items[oxiApp.ItemsListContainer.selectedItem].Name)
		}
	}

	fieldToolbarBottom := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			if oxiApp.FieldsListContainer.currentItem != 0 {
				log.Println("Create field for item #", oxiApp.FieldsListContainer.currentItem)
				oxiApp.Window.SetContent(oxiApp.addFieldForm(oxiApp.FieldsListContainer.currentItem))
				oxiApp.FieldsList.Refresh()
			}

		}),

		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DeleteIcon(), func() {
			log.Println("delete field")
		}),
	)

	fieldContent := container.New(
		layout.NewBorderLayout(oxiApp.ItemNameLabel, fieldToolbarBottom, nil, nil),
		oxiApp.ItemNameLabel, fieldToolbarBottom, oxiApp.FieldsList)

	if oxiApp.ItemsList == nil {
		oxiApp.ItemsList = widget.NewList(oxiApp.ItemsListContainer.GetListSize,
			oxiApp.ItemsListContainer.GetListDesign,
			oxiApp.ItemsListContainer.UpdateItem)
		oxiApp.ItemsList.OnSelected = oxiApp.OnItemSelected
	} else {
		oxiApp.ItemsList.Refresh()
	}

	oxiApp.SplitContainer = container.NewHSplit(oxiApp.ItemsList, fieldContent)
	oxiApp.SplitContainer.SetOffset(oxiApp.App.Preferences().FloatWithFallback(cPrefSplitOffsetKey, cPrefSplitOffset))

	content := container.New(layout.NewBorderLayout(toolbar, searchEntry, nil, nil),
		toolbar, searchEntry, oxiApp.SplitContainer)
	return content
}
