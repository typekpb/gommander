package main

import (
	"github.com/typekpb/gommander/pkg/model"
	"github.com/typekpb/gommander/pkg/ui"
)

func main() {

	modelState := model.NewModelState("/Users/butkovic/all", "/Users/butkovic/Desktop/bs")
	main := ui.NewUiMain(modelState)
	main.Show()
}
