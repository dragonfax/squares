package glt

import "github.com/veandco/go-sdl2/sdl"

var _ Element = &StatelessElement{}

type StatelessElement struct {
	elementData
	childElementData
}

func NewStatelessElement(widget Widget) *StatelessElement {
	se := &StatelessElement{}
	se.widget = widget
	return se
}

func (se StatelessElement) getSize() Size {
	return se.child.getSize()
}

func (se StatelessElement) layout(c Constraints) error {
	return se.child.layout(c)
}

func (se StatelessElement) render(o Offset, r *sdl.Renderer) {
	se.child.render(o, r)
}
