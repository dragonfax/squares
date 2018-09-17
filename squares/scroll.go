package squares

import "github.com/veandco/go-sdl2/sdl"

type CardinalDirection uint8

const (
	CardinalDirectionNone CardinalDirection = iota
	Up
	Down
	Left
	Right
)

type PointerEvent interface {
}

type PointerEventListener interface {
	HandleEvent(PointerEvent) bool
}

var _ PointerEvent = ScrollEvent{}

type ScrollEvent struct {
	Direction CardinalDirection
	Delta     float64
}

var _ ElementWidget = Listener{}
var _ HasChild = Listener{}
var _ Element = &ListenerElement{}
var _ HasChildElement = &ListenerElement{}
var _ PointerEventListener = &ListenerElement{}

type Listener struct {
	Child            Widget
	OnMouseWheelUp   func(PointerEvent) bool
	OnMouseWheelDown func(PointerEvent) bool
}

func (l Listener) getChild() Widget {
	return l.Child
}

func (l Listener) createElement() Element {
	le := &ListenerElement{}
	le.widget = l
	return le
}

type ListenerElement struct {
	elementData
	childElementData
}

func (le *ListenerElement) HandleEvent(event PointerEvent) bool {
	w := le.widget.(Listener)
	if scrollE, ok := event.(ScrollEvent); ok {
		if scrollE.Direction == Up && w.OnMouseWheelUp != nil {
			return w.OnMouseWheelUp(event)
		} else if scrollE.Direction == Down && w.OnMouseWheelDown != nil {
			return w.OnMouseWheelDown(event)
		}
	}
	return false
}

func (le *ListenerElement) layout(c Constraints) error {
	child := le.getChildElement()
	if child == nil {
		return nil
	}
	err := child.layout(c)
	if err != nil {
		return err
	}
	le.size = child.getSize()
	return nil
}

func (le *ListenerElement) render(o Offset, r *sdl.Renderer) {
	child := le.getChildElement()
	if child == nil {
		return
	}
	child.render(o, r)
}
