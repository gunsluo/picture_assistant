package assistant

import (
	"io/ioutil"

	bimg "gopkg.in/h2non/bimg.v1"
)

type FilePictureReader struct {
	path string
}

func NewFilePictureReader(path string) *FilePictureReader {
	return &FilePictureReader{
		path: path,
	}
}

func (reader *FilePictureReader) SetPath(path string) {
	reader.path = path
}

func (reader *FilePictureReader) Read() ([]byte, *PictureInfo, error) {
	buffer, err := ioutil.ReadFile(reader.path)
	if err != nil {
		return nil, nil, err
	}

	size, err := bimg.NewImage(buffer).Size()
	if err != nil {
		return nil, nil, err
	}

	info := &PictureInfo{
		Path: reader.path,
	}
	info.Name = ParseNamePath(reader.path)
	info.Ext = ParseExt(reader.path)
	info.Size = int64(len(buffer))
	info.Width = size.Width
	info.Height = size.Height

	return buffer, info, nil
}
