package squares

type Asset struct {
	File    string
	Package string
}

type BoxFit uint8

const (
	BoxFitCover BoxFit = iota
)

var _ StatelessWidget = Image{}

type Image struct {
	Fit    BoxFit
	Height uint16
}

func NewImageFromAsset(asset Asset, image Image) Image {
	return Image{}
}
