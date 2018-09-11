package glt

var _ StatelessWidget = &Icon{}
var _ StatelessWidget = &IconButton{}
var _ HasChild = &IconButton{}

type VoidCallback func()

type IconData struct {
}

type Icon struct {
	Icon *IconData
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
var IconsCall = &IconData{}
var IconsMessage = &IconData{}
var IconsContactMail = &IconData{}
var IconsEmail = &IconData{}
var IconsLocationOn = &IconData{}
var IconsMap = &IconData{}
var IconsToday = &IconData{}
