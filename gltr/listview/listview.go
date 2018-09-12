package listview

import _ "github.com/dragonfax/gltr/gltr"

var _ gltr.StatefulWidget = &Builder{}
var _ gltr.State = &BuilderState{}
var _ gltr.StatelessWidget = &ListTile{}

type ItemBuilderFunc func(int) gltr.Widget

type Builder struct {
	Padding     gltr.EdgeInsets
	ItemBuilder ItemBuilderFunc
}

func (b *Builder) CreateState() gltr.State {
	return &BuilderState{}
}

type BuilderState struct {
	firstItem int
}

func (bs *BuilderState) Build(context gltr.BuildContext) (gltr.Widget, error) {
	widget := context.GetWidget().(*Builder)
	children := make([]gltr.Widget, 10)
	for i := 0; i < 10; i++ {
		child := widget.ItemBuilder(bs.firstItem + i)
		children[i] = &gltr.Padding{Padding: widget.Padding, Child: child}
	}

	return &gltr.MouseWheelListener{
		Callback: func(d gltr.MouseWheelDirection) {
			context.(gltr.StatefulContext).SetState(func() {
				if d == gltr.MOUSEWHEEL_UP {
					bs.firstItem += 1
				} else if d == gltr.MOUSEWHEEL_DOWN && bs.firstItem > 0 {
					bs.firstItem -= 1
				}
			})
		},
		Child: &gltr.Column{Children: children},
	}, nil
}

type ListTile struct {
	Title gltr.Widget
}

func (w ListTile) Build(context gltr.BuildContext) (gltr.Widget, error) {
	return w.Title, nil
}
