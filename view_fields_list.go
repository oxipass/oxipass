package main

import (
	"github.com/oxipass/oxilib"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type FieldsContainer struct {
	fields      []oxilib.BSField
	currentItem int64
}

func (fc *FieldsContainer) GetListSize() int {
	return len(fc.fields)
}

func (fc *FieldsContainer) GetListDesign() fyne.CanvasObject {

	labelValue := widget.NewLabel("")
	labelName := widget.NewLabelWithStyle("", fyne.TextAlignTrailing, fyne.TextStyle{Italic: true})
	vBox := container.NewVBox(
		labelValue,
		labelName,
	)

	return vBox
}

func (fc *FieldsContainer) UpdateItem(id int, item fyne.CanvasObject) {
	labelValue := item.(*fyne.Container).Objects[0].(*widget.Label)
	labelValue.SetText(fc.fields[id].Value)
	labelName := item.(*fyne.Container).Objects[1].(*widget.Label)
	labelName.SetText(fc.fields[id].Name)
}

func (fc *FieldsContainer) OnSelected(id int) {

	log.Printf("Field #%d '%s' selected\n", id, fc.fields[id].Name)
}

func CreateFieldsList() (listCont *FieldsContainer) {
	listCont = &FieldsContainer{}
	return listCont
}

func UpdateFieldsList(fc *FieldsContainer, itemId int64) (err error) {

	fc.fields, err = oxilib.GetInstance().ReadFieldsByItemID(itemId)
	if err != nil {
		return err
	}
	fc.currentItem = itemId
	return nil
}
