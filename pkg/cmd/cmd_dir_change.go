package cmd

import {
	"github.com/typekpb/gommander/model"

	"os"
}


type DirChange struct {
	state *model.ModelState
	dir string
}

func NewDirChange(state *model.ModelState, dir string) (*cmd.DirChange, error) {
	return &DirChange(state, dir)
}

func (s *DirChange) triggerEvent() {

}

