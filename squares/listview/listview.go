package listview

import (
	"github.com/dragonfax/squares/squares"
)

var _ squares.StatefulWidget = &Builder{}
var _ squares.State = &BuilderState{}
var _ squares.StatelessWidget = &ListTile{}

type ItemBuilderFunc func(int) squares.Widget

type Builder struct {
	Padding     squares.EdgeInsets
	ItemBuilder ItemBuilderFunc
}

func (b *Builder) CreateState() squares.State {
	return &BuilderState{}
}

type BuilderState struct {
	firstItem int
}

func (bs *BuilderState) Build(context squares.StatefulContext) (squares.Widget, error) {
	widget := context.GetWidget().(*Builder)
	children := make([]squares.Widget, 10)
	for i := 0; i < 10; i++ {
		child := widget.ItemBuilder(bs.firstItem + i)
		children[i] = &squares.Padding{Padding: widget.Padding, Child: child}
	}

	return &squares.MouseWheelListener{
		Callback: func(d squares.MouseWheelDirection) {
			context.(squares.StatefulContext).SetState(func() {
				if d == squares.MOUSEWHEEL_UP {
					bs.firstItem += 1
				} else if d == squares.MOUSEWHEEL_DOWN && bs.firstItem > 0 {
					bs.firstItem -= 1
				}
			})
		},
		Child: &squares.Column{Children: children},
	}, nil
}

type ListTile struct {
	Title squares.Widget
}

func (w ListTile) Build(context squares.StatelessContext) (squares.Widget, error) {
	return w.Title, nil
}
