package assistant

var (
	NullPictureConvert PictureConverter
)

type PictureConverter interface {
	Convert([]byte, *PictureInfo) ([]byte, error)
}
