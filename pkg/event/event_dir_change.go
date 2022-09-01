package event

type EventDirChange struct {
	dir string
	files []os.FileInfo
	leftPane bool // or right
}


func NewEventDirChange(dir string, leftPane bool) (*event.EventDirChange) {
	return &EventDirChange{
		dir: dir, 
		leftPane: leftPane
	}
}

type EventDirChangeSubscriber struct {
    Channel chan *EventDirChange
}


// func (s *NewEventDirChange) dirContents() ([]os.FileInfo, error) {
// 	var files []os.FileInfo
// 	f, err := os.Open(root)
// 	if err != nil {
// 		return files, err
// 	}
// 	fileInfo, err := f.Readdir(-1)
// 	f.Close()
// 	if err != nil {
// 		return files, err
// 	}

// 	for _, file := range fileInfo {
// 		files = append(files, file)
// 	}
// 	return files, nil
// }