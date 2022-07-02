package main

import (
	"github.com/oxipass/oxilib"
	"log"

	"fyne.io/fyne/v2/dialog"
)

func (oxiApp *OxiApp) logoutTapped() {
	bsInstance := oxilib.GetInstance()
	err := bsInstance.Lock()
	if err != nil {
		dialog.ShowError(err, oxiApp.Window)
		return
	}
	oxiApp.FieldsList = nil
	oxiApp.FieldsListContainer = nil
	oxiApp.ItemsListContainer = nil
	oxiApp.ItemsList = nil

	oxiApp.Window.SetContent(oxiApp.mainScreen())
	log.Println("tapped func")

}
