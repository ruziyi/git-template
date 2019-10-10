package upload

import (
	"mime/multipart"
	"path/filepath"
)

func IsValidImage(fd *multipart.FileHeader) bool {
	ext := filepath.Ext(fd.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		return false
	}
	t := fd.Header.Get("Content-Type")
	if t != "image/jpeg" && t != "image/png" {
		return false
	}
	return true
}
