package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (oxiApp *OxiApp) lockAction() {
	oxiApp.logoutTapped()
}

func (oxiApp *OxiApp) createMainToolbar() fyne.CanvasObject {
	toolBarItemLock := widget.NewToolbarAction(theme.CancelIcon(), oxiApp.lockAction)
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			oxiApp.Window.SetContent(oxiApp.addItemForm())
		}),

		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.StorageIcon(), func() {
			oxiApp.Window.SetContent(oxiApp.getListScreen())
		}),

		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			oxiApp.Window.SetContent(oxiApp.settingsScreen())
		}),
		widget.NewToolbarSeparator(),
		toolBarItemLock,
	)
}
