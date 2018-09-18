package squares

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

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

var _ ElementWidget = Viewport{}
var _ HasChildren = Viewport{}
var _ Element = &ViewportElement{}
var _ HasChildrenElements = &ViewportElement{}

type Viewport struct {
	Slivers []Widget
}

func (v Viewport) getChildren() []Widget {
	return v.Slivers
}

func (v Viewport) createElement() Element {
	ve := &ViewportElement{}
	ve.widget = v
	return ve
}

type ViewportElement struct {
	InternalOffset Offset
	InternalSize   Size
	elementData
	childrenElementsData
}

func (ve *ViewportElement) layout(c Constraints) error {
	// decide viewport size
	// the viewport itself should fill the parent.
	ve.size = c.constrain(Size{Width: math.Inf(1), Height: math.Inf(1)})

	// start with viewport size.
	// keep width constraints, toss height constraints.
	childConstraints := Constraints{
		minWidth:  0,
		maxWidth:  ve.size.Width,
		minHeight: 0,
		maxHeight: math.Inf(1),
	}

	lastY := 0.0
	for _, child := range ve.children {
		child.layout(childConstraints)
		child.setOffset(Offset{x: 0, y: lastY})
		lastY += child.getSize().Height
	}
	ve.InternalSize = Size{Width: ve.size.Width, Height: lastY}

	return nil
}

func (ve *ViewportElement) render(o Offset, r *sdl.Renderer) {

	// skip children until we find one in the visible viewport.

	for _, child := range ve.children {
		childSize := child.getSize()
		childOffset := child.getOffset()
		if ve.InternalOffset.y <= childOffset.y+childSize.Height && ve.InternalOffset.y+ve.size.Height > childOffset.y {
			no := Offset{x: 0, y: ve.InternalOffset.y}
			child.render(childOffset.Sub(no).Add(o), r)
		}
	}
}

func (ve *ViewportElement) ScrollUp(delta float64) {
	ve.InternalOffset.y += delta
	if ve.InternalOffset.y > ve.InternalSize.Height-ve.size.Height {
		ve.InternalOffset.y = ve.InternalSize.Height - ve.size.Height
	}
}

func (ve *ViewportElement) ScrollDown(delta float64) {
	ve.InternalOffset.y -= delta

	if ve.InternalOffset.y < 0 {
		ve.InternalOffset.y = 0
	}
}

var _ StatelessWidget = SliverList{}
var _ HasChildren = SliverList{}

type SliverList struct {
	Delegate SliverChildListDelegate
}

type SliverChildListDelegate struct {
	Children []Widget
}

func (d SliverList) Build(context StatelessContext) (Widget, error) {
	return &Column{Children: d.getChildren()}, nil
}

func (d SliverList) getChildren() []Widget {
	return d.Delegate.Children
}

var _ StatelessWidget = CustomScrollView{}

type CustomScrollView struct {
	Slivers []Widget
}

func (d CustomScrollView) Build(context StatelessContext) (Widget, error) {
	return Listener{
		Child: Viewport{Slivers: d.Slivers},
		OnMouseWheelUp: func(event PointerEvent) bool {
			context.getElement().getChildElement().(*ListenerElement).getChildElement().(*ViewportElement).ScrollUp(event.(ScrollEvent).Delta * 2)
			return true
		},
		OnMouseWheelDown: func(event PointerEvent) bool {
			context.getElement().getChildElement().(*ListenerElement).getChildElement().(*ViewportElement).ScrollDown(event.(ScrollEvent).Delta * 2)
			return true
		},
	}, nil
}

type SliverAppBar struct {
	ExpandedHeight float64
	Pinned         bool
	Floating       bool
	Snap           bool
	Actions        []Widget
	FlexibleSpace  FlexibleSpaceBar // Widget
}

func (d SliverAppBar) Build(context StatelessContext) (Widget, error) {
	return Stack{
		Children: []Widget{
			d.FlexibleSpace,
			Align{
				Alignment: AlignmentUpperRight,
				Child: Row{
					Children: d.Actions,
				},
			},
		},
	}, nil
}

type FlexibleSpaceBar struct {
	Title      Widget
	Background Widget
}

func (d FlexibleSpaceBar) Build(context StatelessContext) (Widget, error) {
	return Stack{
		Children: []Widget{
			d.Background,
			Align{
				Alignment: AlignmentBottomLeft,
				Child:     d.Title,
			},
		},
	}, nil
}
