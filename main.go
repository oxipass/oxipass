package main

import (
	"fyne.io/fyne/v2"
	"log"
	"os"
	"os/user"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/bykovme/gotrans"
	"github.com/oxipass/oxilib"
)

func main() {
	errLocales := gotrans.InitLocales("langs")
	if errLocales != nil {
		log.Fatal(errLocales)
	}
	errLocales = gotrans.SetDefaultLocale("ru")
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	appFolder := usr.HomeDir + "/" + cAppFolder
	dbFile := appFolder + "/" + cDbFile
	if _, err := os.Stat(appFolder); os.IsNotExist(err) {
		err := os.MkdirAll(appFolder, os.ModePerm)
		if err != nil {
			log.Fatal("Initiation error: " + err.Error())
			return
		}
	}
	oxiInstance := oxilib.GetInstance()
	log.Println("BSApp is initiating ")
	err = oxiInstance.Open(dbFile)
	if err != nil {
		log.Fatal("Initiation error: " + err.Error())
		return
	}
	defer func() {
		log.Println("BSApp stopped, closing storage")
		err := oxiInstance.Close()
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("Storage is closed")
	}()

	appInstance := app.NewWithID(cAppId)
	appInstance.Settings().SetTheme(theme.DefaultTheme())
	// appInstance.Settings().SetTheme(theme.LightTheme())

	window := appInstance.NewWindow(cAppName)
	winWidth := float32(appInstance.Preferences().IntWithFallback(cPrefWidth, cPrefDefWidth))
	winHeight := float32(appInstance.Preferences().IntWithFallback(cPrefHeight, cPrefDefHeight))
	window.Resize(fyne.Size{Width: winWidth, Height: winHeight})
	//window.SetMainMenu()

	bsapp := &OxiApp{}
	bsapp.Init()
	bsapp.Window = window
	bsapp.App = appInstance

	window.SetMaster()
	window.SetContent(bsapp.mainScreen())
	window.SetOnClosed(bsapp.OnClose)
	window.CenterOnScreen()
	window.ShowAndRun()
}
