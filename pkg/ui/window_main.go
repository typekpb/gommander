package ui

import (

	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/typekpb/gommander/pkg/model"
)

type UiMain struct {
	modelState *model.ModelState
}

func NewUiMain(modelState *model.ModelState) *UiMain {
	return &UiMain{
		modelState: modelState,
	}
}

func (m *UiMain) Show() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")

	top := canvas.NewText("top bar", color.White)

	bottom := newUiContainerBottom(m.modelState).build()
	center := newUiContainerCenter(m.modelState).build()

	content := container.New(layout.NewBorderLayout(top, bottom, nil, nil), top, bottom, center)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()

}
