package glt

import "github.com/veandco/go-sdl2/sdl"

var _ ElementWidget = &Text{}
var _ Element = &TextElement{}

const CHARACTER_WIDTH = 10
const CHARACTER_HEIGHT = 10

type Text struct {
	Text string
}

func (t *Text) createElement() Element {
	return &TextElement{widget: t}
}

type TextElement struct {
	widget *Text
	sizeData
	parentData
}

func (t *TextElement) GetWidget() Widget {
	return t.widget
}

func (t *TextElement) layout(c Constraints) error {
	cWidth := len(t.widget.Text) * CHARACTER_WIDTH
	cHeight := 1 * CHARACTER_HEIGHT

	t.size = Size{width: MaxUint16(uint16(cWidth), c.minWidth), height: MaxUint16(uint16(cHeight), c.minHeight)}

	return nil
}

func (t *TextElement) render(offset Offset, renderer *sdl.Renderer) {
	ux := int32(offset.x)
	uy := int32(offset.y)
	surface, err := font.RenderUTF8Blended(t.widget.Text, sdl.Color{R: 200, G: 200, B: 200, A: 255})
	if err != nil {
		panic(err)
	}
	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}
	renderer.Copy(texture, nil, &sdl.Rect{X: ux, Y: uy, W: surface.W, H: surface.H})
}
