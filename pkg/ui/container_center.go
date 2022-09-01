package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/typekpb/gommander/pkg/io"
	"github.com/typekpb/gommander/pkg/model"
)

type containerCenter struct {
	modelState *model.ModelState
}

func newUiContainerCenter(modelState *model.ModelState) *containerCenter {
	return &containerCenter{
		modelState: modelState,
	}
}


func (c *containerCenter) initFilesPane(path string) *widget.List {
	files, err := io.OSReadDir(path)
	if err != nil {
		panic(err)
	}

	// TODO table way to show
	// leftList := widget.NewTable(
	// 	func() (int, int) {
	// 		return len(data), len(data[0])
	// 	},
	// 	func() fyne.CanvasObject {
	// 		return widget.NewLabel("wide content")
	// 	},
	// 	func(i widget.TableCellID, o fyne.CanvasObject) {
	// 		o.(*widget.Label).SetText(data[i.Row][i.Col])
	// 	})

	list := widget.NewList(
		func() int {
			return len(files)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			var name string

			if files[i].IsDir() {
				name = fmt.Sprintf("[%s]", files[i].Name())
			} else {
				name = files[i].Name()
			}

			o.(*widget.Label).SetText(name)
		})
	return list
}

func (c *containerCenter) build() *container.Split {
	
	left := c.initFilesPane(c.modelState.LeftPath())
	right := c.initFilesPane(c.modelState.RightPath())

	// Setup Content
	center := container.NewHSplit(
		left,
		right,
	)

	return center
}
