package glt

type VoidCallback func()

type IconData struct {
}

type Icon struct {
	Icon *IconData
}

type Color struct {
}

type IconButton struct {
	Icon      *Icon
	Tooltip   string
	OnPressed VoidCallback
}

var ColorsIndigo = Color{}

type Brightness uint8

const (
	BrightnessLight Brightness = iota
)

var IconsCreate = &IconData{}
