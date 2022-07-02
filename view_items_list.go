package main

import (
	"errors"
	"github.com/oxipass/oxilib"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ItemsContainer struct {
	items        []oxilib.BSItem
	selectedItem int64
}

func (lc *ItemsContainer) GetListSize() int {
	if lc == nil {
		return 0
	}
	return len(lc.items)
}

func (lc *ItemsContainer) GetListDesign() fyne.CanvasObject {
	hBox := container.NewHBox(
		widget.NewIcon(theme.DocumentIcon()),
		container.NewVBox(
			widget.NewLabel("item name"),
			widget.NewLabel("tags"),
			//itemImage,
		))
	return hBox
}

func (lc *ItemsContainer) UpdateItem(id int, item fyne.CanvasObject) {
	if lc == nil {
		return
	}
	icon := item.(*fyne.Container).Objects[0].(*widget.Icon)
	icon.SetResource(GetIconResource(lc.items[id].Icon))
	labelValue := item.(*fyne.Container).Objects[1].(*fyne.Container).Objects[0].(*widget.Label)
	labelValue.SetText(lc.items[id].Name)
	tags := ""
	for _, tag := range lc.items[id].Tags {
		if tags == "" {
			tags = tag.Name
		} else {
			tags += ", " + tag.Name
		}
	}
}

func (oxiApp *OxiApp) OnItemSelected(id int) {
	oxiApp.ItemsListContainer.selectedItem = oxiApp.ItemsListContainer.items[id].ID
	log.Printf("Item #%d '%s' selected\n",
		oxiApp.ItemsListContainer.selectedItem,
		oxiApp.ItemsListContainer.items[id].Name)
	_ = UpdateFieldsList(oxiApp.FieldsListContainer, oxiApp.ItemsListContainer.items[id].ID)
	oxiApp.FieldsList.Refresh()
	oxiApp.ItemNameLabel.SetText(oxiApp.ItemsListContainer.items[id].Name)

}

func UpdateItemsList(listCont *ItemsContainer) (err error) {
	if listCont == nil {
		return errors.New("container is empty")
	}
	listCont.items, err = oxilib.GetInstance().ReadAllItems(true, false)
	if err != nil {
		return err
	}
	listCont.selectedItem = 0
	return nil
}

func CreateItemsList() (listCont *ItemsContainer, err error) {
	listCont = &ItemsContainer{}
	listCont.items, err = oxilib.GetInstance().ReadAllItems(true, false)
	if err != nil {
		return nil, err
	}
	listCont.selectedItem = 0
	return listCont, nil
}
