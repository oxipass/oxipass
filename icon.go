package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

func GetImagePath(iconName string) string {
	parts := strings.Split(iconName, " ")
	if len(parts) == 2 {
		var iconPath string
		switch parts[0] {
		case "fab":
			iconPath = "brands/"
		case "fas":
			iconPath = "solid/"
		default:
			return ""
		}
		actualName := strings.Replace(parts[1], "fa-", "", -1)
		fullName := "/Users/bykov/BykovSoft/bs/bsapp/icons/" + iconPath + actualName + ".svg"
		return fullName
	}
	return ""
}

func GetIconResource(iconName string) (icon fyne.Resource) {
	icon = theme.DocumentIcon()
	iconResource, err := fyne.LoadResourceFromPath(GetImagePath(iconName))
	if err == nil {
		return iconResource
	}
	return theme.DocumentIcon()
}
