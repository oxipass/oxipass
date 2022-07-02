package main

import (
	"errors"
	"fyne.io/fyne/v2"
	"github.com/oxipass/oxilib"
	"log"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/bykovme/gotrans"
)

func (oxiApp *OxiApp) createLoginForm() fyne.CanvasObject {
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MoveDownIcon(), func() {
			log.Println("Open backup file")
		}),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Help")
			dialog.ShowInformation(cAppName, gotrans.T("version")+" "+cAppVersion, oxiApp.Window)
		}),
	)

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder(gotrans.T("master_password"))

	vBox := container.NewVBox(
		layout.NewSpacer(),
		password,
		widget.NewButtonWithIcon(gotrans.T("unlock"), theme.ConfirmIcon(), func() {
			if password.Text == "" {
				err := errors.New("password cannot be empty")
				dialog.ShowError(err, oxiApp.Window)
				return
			}
			err := oxilib.GetInstance().Unlock(password.Text)
			if err != nil {
				dialog.ShowError(err, oxiApp.Window)
				password.SetText("")
				password.FocusGained()
				return
			}
			oxiApp.Window.SetContent(oxiApp.mainScreen())
		}),
		layout.NewSpacer(),
	)

	content := container.New(layout.NewBorderLayout(toolbar, nil, nil, nil),
		toolbar, vBox)

	password.SetText("")
	password.FocusGained()
	return content
}

func (oxiApp *OxiApp) passwordFieldChanged(key *fyne.KeyEvent) {
	log.Println("Password: " + key.Name)
}
