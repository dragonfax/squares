package glt

type Asset struct {
	File    string
	Package string
}

type BoxFit uint8

const (
	BoxFitCover BoxFit = iota
)

type Image struct {
	Fit    BoxFit
	Height uint16
}

func NewImageFromAsset(asset Asset, image *Image) *Image {
	return nil
}
