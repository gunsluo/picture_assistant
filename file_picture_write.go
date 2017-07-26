package assistant

import (
	"fmt"
	"io/ioutil"
)

type FilePictureWrite struct {
	dir           string
	nameGenerator func(dir string, info *PictureInfo) string
}

func NewFilePictureWrite(dir string) *FilePictureWrite {
	return &FilePictureWrite{
		dir:           dir,
		nameGenerator: defaultNameGenerator,
	}
}

func (write *FilePictureWrite) SetNameGenerator(nameGenerator func(dir string, info *PictureInfo) string) {
	write.nameGenerator = nameGenerator
}

func (write *FilePictureWrite) Write(buffer []byte, info *PictureInfo) error {
	path := write.nameGenerator(write.dir, info)
	err := ioutil.WriteFile(path, buffer, 0644)
	if err != nil {
		return err
	}

	//modify picture info
	info.Path = path
	return nil
}

func defaultNameGenerator(dir string, info *PictureInfo) string {
	return fmt.Sprintf("%s%s-%dx%d.%s", dir, CleanExt(info.Name), info.Width, info.Height, info.Ext)
}

func BulrNameGenerator(dir string, info *PictureInfo) string {
	return fmt.Sprintf("%s%s-%dx%d-bulr.%s", dir, CleanExt(info.Name), info.Width, info.Height, info.Ext)
}
