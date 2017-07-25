package assistant

import (
	"fmt"

	bimg "gopkg.in/h2non/bimg.v1"
)

var (
	DefaultJPGToWEBPConvert = &JPGToWEBPConvert{}
)

type JPGToWEBPConvert struct {
}

func (convert *JPGToWEBPConvert) Convert(buffer []byte, info *PictureInfo) ([]byte, error) {

	newImage, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
	if err != nil {
		return nil, err
	}

	image := bimg.NewImage(newImage)
	if image.Type() != "webp" {
		return nil, fmt.Errorf("The image was converted into webp, failed.")
	}

	size, err := image.Size()
	if err != nil {
		return nil, err
	}

	// modify picture info
	info.Name = ModifyExt(info.Name, PictureTypes.WEBP.String())
	info.Ext = PictureTypes.WEBP
	info.Size = int64(len(newImage))
	info.Width = size.Width
	info.Height = size.Height

	return newImage, nil
}
