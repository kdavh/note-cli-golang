package nconfig

import (
	"os"

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

func NewCfgMock(fs afero.Fs, reporter ReporterInterface, editor EditorInterface) *Config {
	return &Config{
		Searcher:  NewSearcherMock(),
		Editor:    editor,
		Tagline:   "###-tags-:",
		NotesPath: "notes",
		Fs:        fs,
		OsCtrl:    NewOsCtrlMock(),
		Reporter:  reporter,
	}
}

type editorMock struct {
	fs afero.Fs
}

func (e *editorMock) Open(file string, cfg *Config) error {
	return nil
}

func NewEditorMock(fs afero.Fs) *editorMock {
	return &editorMock{
		fs: fs,
	}
}

type searcherMock struct{}

func (s *searcherMock) Notes(namespace string, tagsQuery string, textQuery string, cfg *Config) ([]string, error) {
	return []string{}, nil
}

func NewSearcherMock() *searcherMock {
	return &searcherMock{}
}
