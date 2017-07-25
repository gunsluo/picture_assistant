package assistant

var PictureTypes = struct {
	UNKNOWN PictureType // UNKNOWN represents an unknow image type value.
	JPEG    PictureType // JPEG represents the JPEG image type.
	JPG     PictureType // JPG represents the JPG image type.
	WEBP    PictureType // WEBP represents the WEBP image type.
	PNG     PictureType // PNG represents the PNG image type.
	TIFF    PictureType // TIFF represents the TIFF image type.
	GIF     PictureType // GIF represents the GIF image type.
	PDF     PictureType // PDF represents the PDF type.
	SVG     PictureType // SVG represents the SVG image type.
	MAGICK  PictureType // MAGICK represents the libmagick compatible genetic image type.
}{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
}

type PictureType int

func (pt PictureType) String() string {
	switch pt {
	case PictureTypes.JPEG:
		return "jpeg"
	case PictureTypes.JPG:
		return "jpg"
	case PictureTypes.WEBP:
		return "webp"
	case PictureTypes.PNG:
		return "png"
	case PictureTypes.TIFF:
		return "tiff"
	case PictureTypes.GIF:
		return "gif"
	case PictureTypes.PDF:
		return "pdf"
	case PictureTypes.SVG:
		return "svg"
	case PictureTypes.MAGICK:
		return "magick"
	default:
	}

	return ""
}

type PictureInfo struct {
	Path string      // 全路径
	Name string      // 图片名
	Ext  PictureType // 图片类型
	Size int64       //
	Sha1 string      //
}
