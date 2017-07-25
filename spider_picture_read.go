package assistant

import (
	"io/ioutil"
	"net/http"

	bimg "gopkg.in/h2non/bimg.v1"
)

type SpiderPictureRead struct {
	url string
}

func NewSpiderPictureRead(url string) *SpiderPictureRead {
	return &SpiderPictureRead{
		url: url,
	}
}

func (reader *SpiderPictureRead) Read() ([]byte, *PictureInfo, error) {
	resp, err := http.Get(reader.url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	size, err := bimg.NewImage(buffer).Size()
	if err != nil {
		return nil, nil, err
	}

	info := &PictureInfo{
		Path: reader.url,
	}
	info.Name = ParseNameUrl(reader.url)
	info.Ext = ParseExt(reader.url)
	info.Size = int64(len(buffer))
	info.Width = size.Width
	info.Height = size.Height

	return buffer, info, nil
}
