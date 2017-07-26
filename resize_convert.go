package assistant

import (
	"fmt"

	bimg "gopkg.in/h2non/bimg.v1"
)

var (
	DefaultResizeConvert = &ResizeConvert{}
)

type ResizeConvert struct {
	ratio  bool // 按比例
	width  int  // 图片宽度
	height int  // 图片高度
}

func NewResizeConvert() *ResizeConvert {
	return &ResizeConvert{}
}

func (convert *ResizeConvert) SetRatio(ratio bool) *ResizeConvert {
	convert.ratio = ratio
	return convert
}

func (convert *ResizeConvert) SetWidth(width int) *ResizeConvert {
	convert.width = width
	return convert
}

func (convert *ResizeConvert) SetHeight(height int) *ResizeConvert {
	convert.height = height
	return convert
}

func (convert *ResizeConvert) resize(width, height int) (force bool) {

	if convert.ratio == false {
		// 按指定大小缩放
		force = true
		return
	}

	if convert.width != 0 {
		// 按宽度比例缩放
		convert.height = convert.width * height / width
	} else {
		// 按高度比例缩放
		convert.width = convert.height * width / height
	}

	if width < convert.width || height < convert.height {
		force = true
	}

	return
}

func (convert *ResizeConvert) Convert(buffer []byte, info *PictureInfo) ([]byte, error) {

	image := bimg.NewImage(buffer)
	size, err := image.Size()
	if err != nil {
		return nil, err
	}
	force := convert.resize(size.Width, size.Height)

	var newImage []byte
	if force {
		newImage, err = image.ForceResize(convert.width, convert.height)
	} else {
		newImage, err = image.Resize(convert.width, convert.height)
	}
	if err != nil {
		return nil, err
	}

	nsize, err := bimg.NewImage(newImage).Size()
	if err != nil {
		return nil, err
	}
	if nsize.Width != convert.width || nsize.Height != convert.height {
		return nil, fmt.Errorf("resize image size is invalid")
	}

	// modify picture info
	info.Size = int64(len(newImage))
	info.Width = convert.width
	info.Height = convert.height

	return newImage, nil
}
