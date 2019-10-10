package image

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"github.com/smallnest/rpcx/log"
	"gst/pkg/upload"
	"gst/pkg/util"
	"image"
	"image/jpeg"
	"image/png"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type SaveFn func(ctx *gin.Context, pic *multipart.FileHeader, options ...ResizeOption) (string, error)

type ResizeOption struct {
	Width  uint
	Height uint
}

func ResizeImg(fp string, width uint, height uint) error {
	ext := filepath.Ext(fp)
	file, err := os.OpenFile(fp, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	var img image.Image
	if ext == ".jpg" || ext == ".jpeg" {
		img, err = jpeg.Decode(file)
	} else if ext == ".png" {
		img, err = png.Decode(file)
	} else {
		return errors.New("不支持图片类型")
	}
	if err != nil {
		return err
	}
	m := resize.Resize(width, height, img, resize.Lanczos3)

	out, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer out.Close()

	if ext == ".jpg" || ext == ".jpeg" {
		err = jpeg.Encode(out, m, nil)
	} else {
		err = png.Encode(out, m)
	}
	if err != nil {
		return err
	}
	return nil
}

func SavePicFunc(uploadPath string) SaveFn {
	return func(ctx *gin.Context, pic *multipart.FileHeader, options ...ResizeOption) (string, error) {
		if !upload.IsValidImage(pic) {
			return "", errors.New("错误图片格式")
		}
		filename := time.Now().Format("20060102150405") + fmt.Sprintf("%d", rand.Intn(10000)+1000)
		//file := "/static/uploads/" + filename + filepath.Ext(pic.Filename)
		file := uploadPath + filename + filepath.Ext(pic.Filename)
		p := filepath.Join(util.ExecPath(), file)
		err := ctx.SaveUploadedFile(pic, p)
		if err != nil {
			return "", errors.New("上传失败")
		}
		if len(options) > 0 {
			option := options[0]
			err = ResizeImg(p, option.Width, option.Height)
			if err != nil {
				log.Warn("压缩图片错误", err)
			}
		}
		return file, nil
	}
}
