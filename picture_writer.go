package assistant

type PictureWriter interface {
	Write([]byte, *PictureInfo) error
}
