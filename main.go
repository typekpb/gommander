package main

import (
	"fmt"
	"log"

	"image/color"

	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var data = [][]string{[]string{"bin", "opt"},
	[]string{"usr", "tmp"}}

var listData = []string{"a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list"}

func OSReadDir(root string) ([]os.FileInfo, error) {
	var files []os.FileInfo
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file)
	}
	return files, nil
}

func main() {
	var (
		files []os.FileInfo
		err   error
	)

	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")

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

	top := canvas.NewText("top bar", color.White)

	leftList := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	files, err = OSReadDir("/Users/butkovic")
	if err != nil {
		panic(err)
	}

	rightList := widget.NewList(
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

	// Setup Content
	middle := container.NewHSplit(
		leftList,
		// 2nd Section
		rightList,
	)

	content := container.New(layout.NewBorderLayout(top, bottom, nil, nil), top, bottom, middle)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
