package ui

// import (
// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/widget"
// )

// // files panel
// type PanelFiles struct {
// 	modelState *model.modelState
// 	leftPane   *widget.BaseWidget
// 	rightPane  *widget.BaseWidget
// }

// func NewPanelFiles(modelState *model.modelState) *PanelFiles {

// 	// uri := os.Getenv("mongo_uri")
// 	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// return mongoStore{Client: client}, nil

// 	// leftList := widget.NewTable(
// 	// 	func() (int, int) {
// 	// 		return len(data), len(data[0])
// 	// 	},
// 	// 	func() fyne.CanvasObject {
// 	// 		return widget.NewLabel("wide content")
// 	// 	},
// 	// 	func(i widget.TableCellID, o fyne.CanvasObject) {
// 	// 		o.(*widget.Label).SetText(data[i.Row][i.Col])
// 	// 	})

// 	// files, err = OSReadDir("/Users/butkovic")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	leftPane = widget.NewList(
// 		func() int {
// 			return 0
// 		},
// 		func() fyne.CanvasObject {
// 			return widget.NewLabel("template")
// 		},
// 		func(i widget.ListItemID, o fyne.CanvasObject) {
// 		})

// 	rightPane = widget.NewList(
// 		func() int {
// 			return 0
// 		},
// 		func() fyne.CanvasObject {
// 			return widget.NewLabel("template")
// 		},
// 		func(i widget.ListItemID, o fyne.CanvasObject) {
// 		})

// 	return &PanelFiles{
// 		modelState: model.modelState,
// 		lefPane:    leftPane,
// 		rightPane:  rightPane}

// }

// func handleEvent(eventDirChange *event.EventDirChange) {
// 	files, err = OSReadDir(eventDirChange.dir)
// 	if err != nil {
// 		panic(err)
// 	}
// }
