package nconfig

import (
	"github.com/spf13/afero"
)

type exitFunc func(int)

type OsCtrl struct {
	Exit exitFunc
}

type Config struct {
	SearchApp    string
	Editor       string
	EditorConfig string
	Tagline      string
	NotesPath    string
	Fs           afero.Fs
	OsCtrl       OsCtrl
}
