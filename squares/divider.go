package squares

import (
	"github.com/veandco/go-sdl2/sdl"
)

var _ ElementWidget = &Divider{}
var _ Element = &DividerElement{}

type Divider struct {
}

func (d *Divider) createElement() Element {
	de := &DividerElement{}
	de.widget = d
	return de
}

type DividerElement struct {
	elementData
}

func (ce *DividerElement) layout(c Constraints) error {

	ce.size = Size{width: c.maxWidth, height: 16}

	return nil
}

func (ce *DividerElement) render(offset Offset, renderer *sdl.Renderer) {

	renderer.SetDrawColor(0x80, 0x80, 0x80, 255)
	ux := int32(offset.x)
	uy := int32(offset.y + ce.size.height/2)
	renderer.DrawLine(ux+0, uy, ux+int32(ce.size.width), uy)
}
