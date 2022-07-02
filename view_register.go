package main

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/oxipass/oxilib"
	"os"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (oxiApp *OxiApp) createRegisterForm() fyne.CanvasObject {
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("New Master Password")
	passwordConfirm := widget.NewPasswordEntry()
	passwordConfirm.SetPlaceHolder("Confirm Master Password")
	form := &widget.Form{
		OnCancel: func() {
			fmt.Println("Cancelled")
			os.Exit(0)
		},
		OnSubmit: func() {
			if password.Text == "" {
				err := errors.New("password cannot be empty")
				dialog.ShowError(err, oxiApp.Window)
				return
			} else if password.Text != passwordConfirm.Text {
				err := errors.New("password should be the same")
				dialog.ShowError(err, oxiApp.Window)
				return
			}
			err := oxilib.GetInstance().SetNewPassword(password.Text, "AES256V1")
			if err != nil {
				dialog.ShowError(err, oxiApp.Window)
				return
			}
			oxiApp.Window.SetContent(oxiApp.mainScreen())
		},
	}

	form.Append("Password", password)
	form.Append("Confirm", passwordConfirm)

	return form
}
