package assistant

type PictureConverter interface {
	Convert([]byte, *PictureInfo) ([]byte, error)
}
