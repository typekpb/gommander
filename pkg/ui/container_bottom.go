package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/typekpb/gommander/pkg/model"
)

type containerBottom struct {
	modelState *model.ModelState
}

func newUiContainerBottom(modelState *model.ModelState) *containerBottom {
	return &containerBottom{
		modelState: modelState,
	}
}

func (m *containerBottom) build() *fyne.Container {

	viewBtn := widget.NewButton("View", func() {
		log.Println("tapped")
	})
	editBtn := widget.NewButton("Edit", func() {
		log.Println("tapped")
	})
	copyBtn := widget.NewButton("Copy", func() {
		log.Println("tapped")
	})
	moveBtn := widget.NewButton("Move", func() {
		log.Println("tapped")
	})
	mkdirBtn := widget.NewButton("Directory", func() {
		log.Println("tapped")
	})
	deleteBtn := widget.NewButton("Delete", func() {
		log.Println("tapped")
	})
	bottom := container.New(layout.NewGridLayout(6), viewBtn, editBtn, copyBtn, moveBtn, mkdirBtn, deleteBtn)
	return bottom
}
