package util

import (
	"os"
	"path/filepath"
)

func ExecPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}
