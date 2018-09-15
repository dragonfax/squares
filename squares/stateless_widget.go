package squares

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
	if se.child != nil {
		return se.child.getSize()
	} else {
		return Size{}
	}
}

func (se StatelessElement) layout(c Constraints) error {
	if se.child != nil {
		return se.child.layout(c)
	} else {
		return nil
	}
}

func (se StatelessElement) render(o Offset, r *sdl.Renderer) {
	if se.child != nil {
		se.child.render(o, r)
	}
}
