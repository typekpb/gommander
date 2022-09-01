package model

// enum Pane -> Left, Right
type Pane int

const (
	Left Pane = iota
	Right
)

type ModelState struct {
	left  *ModelPane
	right *ModelPane

	activePane Pane
}

type ModelPane struct {
	path string
	// cursor    string
	// selection []string
}

func NewModelState(leftPaneDir string, rightPaneDir string) *ModelState {
	left := &ModelPane{
		path: leftPaneDir,
	}
	right := &ModelPane{
		path: rightPaneDir,
	}
	return &ModelState{
		left:       left,
		right:      right,
		activePane: Left,
	}
}

func (s *ModelState) LeftPath() string {
	return s.left.path
}

func (s *ModelState) RightPath() string {
	return s.right.path
}
