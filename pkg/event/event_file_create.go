package event

type EventFileCreate struct {
	fileName string
	dir string
}


func NewEventFileCreate(fileName string, dir string) (*event.EventDirChange) {
	return &EventFileCreate{	
		fileName: fileName, 
		dir: dir
	}	
}