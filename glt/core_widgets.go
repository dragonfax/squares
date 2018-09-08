package glt

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Divider struct {
	sizeData
	parentData
}

func (ce *Divider) layout(c Constraints) error {

	ce.size = Size{width: c.maxWidth, height: 5}

	return nil
}

func (ce *Divider) render(offset Offset, renderer *sdl.Renderer) {

	renderer.SetDrawColor(0x80, 0x80, 0x80, 255)
	ux := int32(offset.x)
	uy := int32(offset.y)
	renderer.DrawLine(ux+0, uy+3, ux+int32(ce.size.width), uy+3)
}

const CHARACTER_WIDTH = 10
const CHARACTER_HEIGHT = 10

type Text struct {
	Text string
	sizeData
	parentData
}

func (t *Text) layout(c Constraints) error {
	cWidth := len(t.Text) * CHARACTER_WIDTH
	cHeight := 1 * CHARACTER_HEIGHT

	t.size = Size{width: MaxUint16(uint16(cWidth), c.minWidth), height: MaxUint16(uint16(cHeight), c.minHeight)}

	return nil
}

func (t *Text) render(offset Offset, renderer *sdl.Renderer) {
	ux := int32(offset.x)
	uy := int32(offset.y)
	surface, err := font.RenderUTF8Blended(t.Text, sdl.Color{R: 200, G: 200, B: 200, A: 255})
	if err != nil {
		panic(err)
	}
	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}
	renderer.Copy(texture, nil, &sdl.Rect{X: ux, Y: uy, W: surface.W, H: surface.H})
}
