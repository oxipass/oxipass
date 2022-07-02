package main

import (
	"fyne.io/fyne/v2"
	"github.com/oxipass/oxilib"
	"log"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TagsContainer struct {
	tags []oxilib.BSTag
}

func (tc *TagsContainer) GetListSize() int {
	if tc == nil {
		return 0
	}
	return len(tc.tags)
}

func (tc *TagsContainer) GetListDesign() fyne.CanvasObject {
	return widget.NewLabel("")
}

func (oxiApp *OxiApp) getListScreen() fyne.CanvasObject {
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),

		widget.NewToolbarAction(theme.ContentClearIcon(), func() {
			oxiApp.Window.SetContent(oxiApp.mainScreen())
		}),
	)
	tagsToolbarBottom := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			log.Println("Creating a new tag")

			wEntryForKey := widget.NewEntry()
			wEntryForKey.SetPlaceHolder("New tag name")
			var items []*widget.FormItem

			items = []*widget.FormItem{ // we can specify items in the constructor
				{Widget: wEntryForKey},
			}

			dialog.ShowForm("Creating a new tag", "Create", "Close", items, func(entered bool) {
				if entered {
					var newTag oxilib.UpdateTagForm
					newTag.Name = wEntryForKey.Text

					bsInst := oxilib.GetInstance()
					resp, err := bsInst.AddNewTag(newTag)
					if err == nil && resp.Status == oxilib.ConstSuccessResponse {
						if oxiApp.TagsListContainer != nil {
							_ = UpdateTagsList(oxiApp.TagsListContainer)
						}
					}
				}
			}, oxiApp.Window)
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DeleteIcon(), func() {
			log.Println("Delete tag")
		}),
	)

	if oxiApp.TagsListContainer == nil {
		oxiApp.TagsListContainer = CreateTagsList()
	} else {
		_ = UpdateTagsList(oxiApp.TagsListContainer)
	}

	if oxiApp.TagsList == nil {
		oxiApp.TagsList = widget.NewList(oxiApp.TagsListContainer.GetListSize,
			oxiApp.TagsListContainer.GetListDesign,
			oxiApp.TagsListContainer.UpdateTag)
		oxiApp.TagsList.OnSelected = oxiApp.TagsListContainer.OnSelected
	}

	headerLabel := widget.NewLabel("Tags management")

	tagsContent := container.New(
		layout.NewBorderLayout(headerLabel, tagsToolbarBottom, nil, nil),
		headerLabel, tagsToolbarBottom, oxiApp.TagsList)
	content := container.New(layout.NewBorderLayout(toolbar, nil, nil, nil),
		toolbar, tagsContent)
	return content

}

func (tc *TagsContainer) UpdateTag(id int, tag fyne.CanvasObject) {
	tagValue := tag.(*widget.Label)
	tagValue.SetText(tc.tags[id].Name)
}

func (tc *TagsContainer) OnSelected(id int) {
	log.Printf("Tag #%d '%s' selected\n", id, tc.tags[id].Name)
}

func CreateTagsList() (listCont *TagsContainer) {
	listCont = &TagsContainer{}
	err := UpdateTagsList(listCont)
	if err != nil {
		log.Println(err.Error())
	}
	return listCont
}

func UpdateTagsList(tc *TagsContainer) (err error) {

	tc.tags, err = oxilib.GetInstance().GetTags()
	if err != nil {
		return err
	}
	return nil
}
