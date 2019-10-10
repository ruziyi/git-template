package ginUtil

import (
	"net/http"
	"os"
)

//禁止访问目录
type JustFilesFilesystem struct {
	fs http.FileSystem
}

func NewJustFilesFilesystem(fs http.FileSystem) JustFilesFilesystem {
	return JustFilesFilesystem{fs: fs}
}

func (fs JustFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

type neuteredReaddirFile struct {
	http.File
}

func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}
