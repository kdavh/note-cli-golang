package nconfig

import (
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

func NewOsCtrlMock() OsCtrl {
	return OsCtrl{
		Exit: func(code int) {
			if code >= 1 {
				panic(code)
			}
		},
		IsExist: os.IsExist,
	}
}

func NewCfgMock(fs afero.Fs, reporter ReporterInterface) *Config {
	return &Config{
		SearchApp:    "ag",
		Editor:       "nvim",
		EditorConfig: filepath.Join("mocks", "editor-config"),
		Tagline:      "###-tags-:",
		NotesPath:    "notes",
		Fs:           fs,
		OsCtrl:       NewOsCtrlMock(),
		Reporter:     reporter,
	}
}
