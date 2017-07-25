package assistant

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/snowflake"
)

var WorkerId = int64(1)

var generator *snowflake.Node

func uuid() (ts int64, err error) {

	if generator == nil {
		if iw, err := snowflake.NewNode(WorkerId); err != nil {
			return 0, err
		} else {
			generator = iw
		}
	}

	return generator.Generate().Int64(), nil
}

func AutoName() (string, error) {
	ts, err := uuid()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", ts), nil
}

func ParseNameUrl(url string) string {
	i := strings.LastIndexByte(url, '/')
	if i < 0 {
		i = 0
	} else {
		i++
	}

	return url[i:]
}

func ParseNamePath(path string) string {
	i := strings.LastIndexByte(path, '/')
	if i < 0 {
		i = 0
	} else {
		i++
	}

	return path[i:]
}

func ParseExt(path string) PictureType {

	j := strings.LastIndexByte(path, '.')
	if j < 0 {
		return PictureTypes.UNKNOWN
	}

	var ext PictureType
	typ := path[j+1:]
	switch typ {
	case "jpeg":
		ext = PictureTypes.JPEG
	case "jpg":
		ext = PictureTypes.JPG
	case "webp":
		ext = PictureTypes.WEBP
	case "png":
		ext = PictureTypes.PNG
	case "tiff":
		ext = PictureTypes.TIFF
	case "gif":
		ext = PictureTypes.GIF
	case "pdf":
		ext = PictureTypes.PDF
	case "svg":
		ext = PictureTypes.SVG
	case "magick":
		ext = PictureTypes.MAGICK
	default:
		ext = PictureTypes.UNKNOWN
	}

	return ext
}

func ModifyExt(path, ext string) string {

	i := strings.LastIndexByte(path, '.')
	if i < 0 {
		return fmt.Sprintf("%s.%s", path, ext)
	}

	return fmt.Sprintf("%s.%s", path[0:i], ext)
}
