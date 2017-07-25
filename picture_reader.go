package assistant

type PictureReader interface {
	Read() ([]byte, *PictureInfo, error)
}
