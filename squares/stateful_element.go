package squares

import "github.com/veandco/go-sdl2/sdl"

type StatefulElement struct {
	elementData
	state State
	childElementData
	built    bool
	rendered bool
}

var _ Element = &StatefulElement{}
var _ StatefulContext = &StatefulElement{}

func NewStatefulElement(widget Widget, state State) *StatefulElement {
	se := &StatefulElement{}
	se.widget = widget
	se.state = state
	return se
}

func (se *StatefulElement) getSize() Size {
	if se.child != nil {
		return se.child.getSize()
	} else {
		return Size{}
	}
}

func (se *StatefulElement) layout(c Constraints) error {
	if se.child != nil {
		return se.child.layout(c)
	} else {
		return nil
	}
}

func (se *StatefulElement) updateWidget(widget Widget) {
	se.elementData.updateWidget(widget)
}

func (se *StatefulElement) render(o Offset, r *sdl.Renderer) {
	if se.child != nil {
		se.child.render(o, r)
	}
}

func (se *StatefulElement) SetState(callback SetStateFunc) {
	callback()
	se.built = false
}

func (se *StatefulElement) GetState() State {
	return se.state
}
