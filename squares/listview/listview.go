package listview

import (
	. "github.com/dragonfax/squares/squares"
)

var _ StatefulWidget = Builder{}
var _ State = &BuilderState{}
var _ StatelessWidget = ListTile{}

type ItemBuilderFunc func(int) Widget

type Builder struct {
	Padding     EdgeInsets
	ItemBuilder ItemBuilderFunc
}

func (b Builder) CreateState() State {
	return &BuilderState{}
}

type BuilderState struct {
	firstItem int
}

func (bs BuilderState) Build(context StatefulContext) (Widget, error) {
	widget := context.GetWidget().(Builder)
	children := make([]Widget, 10)
	for i := 0; i < 10; i++ {
		child := widget.ItemBuilder(bs.firstItem + i)
		children[i] = &Padding{Padding: widget.Padding, Child: child}
	}

	return Listener{
		OnMouseWheelDown: func(d PointerEvent) bool {
			context.SetState(func() {
				if bs.firstItem > 0 {
					bs.firstItem -= 1
				}
			})
			return true
		},
		OnMouseWheelUp: func(d PointerEvent) bool {
			context.SetState(func() {
				bs.firstItem += 1
			})
			return true
		},
		Child: &Column{Children: children},
	}, nil
}

type ListTile struct {
	Title Widget
}

func (w ListTile) Build(context StatelessContext) (Widget, error) {
	return w.Title, nil
}
