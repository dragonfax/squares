package glt

import "github.com/veandco/go-sdl2/sdl"

type StatefulElement struct {
	elementData
	state State
	childElementData
}

var _ Element = &StatefulElement{}

func NewStatefulElement(widget Widget, state State) *StatefulElement {
	se := &StatefulElement{}
	se.widget = widget
	se.state = state
	return se
}

func (se StatefulElement) getSize() Size {
	return se.child.getSize()
}

func (se StatefulElement) layout(c Constraints) error {
	return se.child.layout(c)
}

func (se StatefulElement) render(o Offset, r *sdl.Renderer) {
	se.child.render(o, r)
}
