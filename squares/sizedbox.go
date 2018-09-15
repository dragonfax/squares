package squares

import "github.com/veandco/go-sdl2/sdl"

var _ ElementWidget = &SizedBox{}
var _ HasChild = &SizedBox{}
var _ Element = &SizedBoxElement{}
var _ HasChildElement = &SizedBoxElement{}

type SizedBox struct {
	Size  Size
	Child Widget
}

func (sb *SizedBox) createElement() Element {
	sbe := &SizedBoxElement{}
	sbe.widget = sb
	return sbe
}

type SizedBoxElement struct {
	widget Widget
	elementData
	childElementData
}

func (sbe *SizedBoxElement) layout(constraints Constraints) error {
	widget := sbe.widget.(*SizedBox)

	if sbe.child != nil {
		sbe.child.layout(ConstraintsAbsolute(widget.Size.Width, widget.Size.Height))
	}

	sbe.size = widget.Size

	return nil
}

func (sbe *SizedBoxElement) render(offset Offset, renderer *sdl.Renderer) {
	if sbe.child != nil {
		sbe.child.render(offset, renderer)
	}
}
