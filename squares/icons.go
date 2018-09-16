package squares

import (
	"path"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var iconFont *ttf.Font

func initIcons() {
	_, file, _, _ := runtime.Caller(0)
	var err error
	iconFont, err = ttf.OpenFont(path.Dir(file)+"/fonts/MaterialIcons-Regular.ttf", 24)
	if err != nil {
		panic(err)
	}
}

var _ ElementWidget = &Icon{}
var _ Element = &IconElement{}
var _ StatelessWidget = &IconButton{}
var _ HasChild = &IconButton{}

type VoidCallback func()

type IconData struct {
	CodePoint rune
}

type Icon struct {
	Icon *IconData
}

func (i *Icon) createElement() Element {
	ie := &IconElement{}
	ie.widget = i
	ie.size = Size{24, 24}
	return ie
}

type IconElement struct {
	elementData
}

func (ie *IconElement) layout(c Constraints) error {
	return nil
}

func (t *IconElement) render(offset Offset, renderer *sdl.Renderer) {
	ux := int32(offset.x)
	uy := int32(offset.y)
	surface, err := iconFont.RenderUTF8Blended(string(t.widget.(*Icon).Icon.CodePoint), sdl.Color{R: 200, G: 200, B: 200, A: 255})
	if err != nil {
		panic(err)
	}
	defer surface.Free()
	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()
	renderer.Copy(texture, nil, &sdl.Rect{X: ux, Y: uy, W: surface.W, H: surface.H})
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

// format for the icon data lines
// var IconsCreate = &IconData{'\ue150'}
