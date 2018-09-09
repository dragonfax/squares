package glt

import "github.com/veandco/go-sdl2/sdl"

var _ Element = &StatelessElement{}

type StatelessElement struct {
	widget Widget
	child  Element
}

func (se *StatelessElement) GetWidget() Widget {
	return se.widget
}

func (se StatelessElement) getSize() Size {
	return se.child.getSize()
}

func (se StatelessElement) getParentData() *parentData {
	return se.child.getParentData()
}

func (se StatelessElement) layout(c Constraints) error {
	return se.child.layout(c)
}

func (se StatelessElement) render(o Offset, r *sdl.Renderer) {
	se.child.render(o, r)
}
