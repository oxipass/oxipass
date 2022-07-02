package main

import (
	"fyne.io/fyne/v2"
	"log"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/bykovme/gotrans"
)

func (oxiApp *OxiApp) settingsScreen() fyne.CanvasObject {
	securityTab := container.NewVBox(
		widget.NewButtonWithIcon("Change password", theme.ConfirmIcon(), func() {}),
		layout.NewSpacer())

	aboutTab := oxiApp.buildAboutView()

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon(gotrans.T("general"), theme.HomeIcon(), widget.NewLabel("Hello")),
		container.NewTabItemWithIcon(gotrans.T("security"), theme.ComputerIcon(), securityTab),
		container.NewTabItemWithIcon(gotrans.T("about"), theme.HelpIcon(), aboutTab))

	//widget.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab"))

	tabs.SetTabLocation(container.TabLocationTop)

	vBox := container.NewVBox(
		widget.NewLabelWithStyle("Settings", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		tabs,
		layout.NewSpacer(),

		widget.NewButtonWithIcon("Close", theme.CancelIcon(), func() {
			oxiApp.Window.SetContent(oxiApp.mainScreen())
			log.Println("Close tapped")
		}),
	)

	return vBox
}
